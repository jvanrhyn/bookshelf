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

func getBookById(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			slog.Error(err.Error(), "idStr", idStr)
			return err
		}
		err = ctx.Status(fiber.StatusOK).JSON(id)
		if err != nil {
			slog.Error(err.Error())
		}
	}
	return nil
}
