package main

import (
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/jvanrhyn/bookshelf/internal/api"
	"github.com/jvanrhyn/bookshelf/internal/controller"
	"github.com/jvanrhyn/bookshelf/internal/database"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

/*
init is a special function; it's called automatically when the package is imported.
In this function, the environment variables are loaded from a .env file using the godotenv package.
If the loading fails, it checks for a custom .env file path using the GetEnvFilePath function from the api package.
If the custom loading fails too, it panics and the application will stop, printing the error message.
In regular use, the function does not require any arguments or returns any values.
Panic in Go is equivalent to exceptions in other languages.
*/
func init() {

	err := godotenv.Load()
	if err != nil {
		envPath := api.GetEnvFilePath()
		err = godotenv.Load(envPath)
		if err != nil {
			panic("Error loading .env file : " + err.Error())
		}
	}
}

func main() {

	logTo := os.Getenv("LOG_TO")
	if logTo == "" {
		logTo = "console"
	} else {
		logTo = strings.ToLower(logTo)
	}

	var opts *slog.HandlerOptions
	switch logTo {
	case "console":
		w := os.Stderr
		handler := log.New(w)
		handler.SetLevel(log.DebugLevel)
		handler.SetTimeFormat(time.Kitchen)
		handler.SetReportTimestamp(true)

		slog.SetDefault(slog.New(
			handler))
	case "json":
		opts = &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}

		logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
		slog.SetDefault(logger)

	default:
		opts = &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}

		logger := slog.NewTextHandler(os.Stderr, opts)
		slog.SetDefault(slog.New(logger))
		slog.Warn("No value specified for the environment value LOG_TO")
	}

	slog.Info("Starting the application")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := database.InitDatabase()
		if err != nil {
			slog.Error(err.Error())
			close(stop)
		}

		err = controller.StartAndServe()
		if err != nil {
			slog.Error(err.Error())
			close(stop)
		}
	}()

	<-stop

	controller.Shutdown()
}
