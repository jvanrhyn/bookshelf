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
	"syscall"
	"time"
)

func init() {

	// Read Configuration data from the .env file in the project
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

	w := os.Stderr
	handler := log.New(w)
	handler.SetLevel(log.DebugLevel)
	handler.SetTimeFormat(time.Kitchen)
	handler.SetReportTimestamp(true)

	slog.SetDefault(slog.New(
		handler))
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
