package controller

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
)

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
