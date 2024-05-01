// Package controller contains the main entry point for the application.
package controller

import (
	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
	"log/slog"
	"os"
)

// StartAndServe starts the application server and listens for incoming requests.
// It initializes the logger, sets the port from the environment variable,
// logs a message about starting the server, creates a new Fiber app,
// sets up the logger middleware, creates the groups, and finally starts listening
// on the specified port.
//
// Parameters:
//
//	None
//
// Returns:
//
//	error: If there is an error starting the server, it will be returned.
//	nil: If the server starts successfully, nil will be returned.
func StartAndServe() error {
	logger := slog.Default()
	port := os.Getenv("PORT")

	go slog.Info("Starting server on port", "port", port)

	app := fiber.New()
	app.Use(slogfiber.New(logger))

	err := createGroups(app)
	if err != nil {
		slog.Error("Error creating groups", "error", err)
	}

	err = app.Listen(":" + port)
	if err != nil {
		go slog.Error("Error starting server", "error", err)
		return err
	}

	return nil
}

// createGroups creates the necessary groups for the application.
// It registers the book endpoints using the registerBookEndPoints function.
// If there is an error registering the book endpoints, it logs an error message.
//
// Parameters:
//
//	app *fiber.App: The Fiber application instance.
//
// Returns:
//
//	nil: If the groups are created successfully, nil will be returned.
//	error: If there is an error registering the book endpoints, it will be returned.
func createGroups(app *fiber.App) error {
	err := registerBookEndPoints(app)
	if err != nil {
		slog.Error("Error registering book endpoints", "error", err)
	}
	return nil
}

func Shutdown() {
	slog.Info("Server is shutting down. Have a great day")
}
