package controller

import "github.com/gofiber/fiber/v2"

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
	return nil
}
