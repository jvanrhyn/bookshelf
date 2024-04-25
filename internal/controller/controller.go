package controller

import (
	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
	"log/slog"
	"os"
)

func StartAndServe() error {
	logger := slog.Default()
	port := os.Getenv("PORT")

	go slog.Info("Starting server on port", "port", port)

	app := fiber.New()
	app.Use(slogfiber.New(logger))

	err := app.Listen(":" + port)
	if err != nil {
		go slog.Error("Error starting server", "error", err)
		return err
	}

	return nil
}

func Shutdown() {
	slog.Info("Server is shutting down. Have a great day")
}
