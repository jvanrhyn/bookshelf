package controller

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
)

// registerBookEndPoints sets up the routes for the book-related endpoints.
//
// This function takes a fiber.App object as argument and uses it
// to define the group of routes prefixed with "books/".
//
// The following endpoints are registered in this function:
// - GET "/:id" that maps to the getBookById function,
// - GET "/isbn/:isbn" that maps to the getBookByISBN function,
// - POST "/" that maps to the createBook function,
// - PUT "/:id" that maps to the updateBook function,
// - DELETE "/:id" that maps to the deleteBook function.
//
// It does not return any values.
func registerBookEndPoints(app *fiber.App) {
	books := app.Group("books/")

	books.Get("/:id", getBookById)
	books.Get("/isbn/:isbn", getBookByISBN)
	books.Post("", createBook)
	books.Put("/:id", updateBook)
	books.Delete("/:id", deleteBook)
}

func deleteBook(ctx *fiber.Ctx) error {

	id, err := getIdFromContext(ctx)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info("Deleting book", "id", id)
	return nil
}

func updateBook(ctx *fiber.Ctx) error {
	return nil
}

func createBook(ctx *fiber.Ctx) error {
	return nil
}

func getBookByISBN(ctx *fiber.Ctx) error {
	return nil
}

// getBookById is a function that is used as a handler for the GET "/:id" route.
// It attempts to get a book's information based on its ID from the request parameters.
// If the ID is not empty and can be parsed to an integer, the function responds with the ID in JSON format.
// If an error occurs during the parsing, the error is logged and returned.
// If an error occurs while responding with JSON, the error is logged.
// This function is expected to be used in a context where the router makes sure the ID parameter exists.
// It returns an error that represents how the function execution went.
//
// The function takes a *fiber.Ctx object as its argument which represents the context of the request.
func getBookById(ctx *fiber.Ctx) error {

	id, err := getIdFromContext(ctx)
	err = ctx.Status(fiber.StatusOK).JSON(id)
	if err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func getIdFromContext(ctx *fiber.Ctx) (int, error) {
	idStr := ctx.Params("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			slog.Error(err.Error(), "idStr", idStr)
			return 0, err
		}
		return id, nil
	}
	return 0, nil
}
