package controller

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jvanrhyn/bookshelf/internal/database"
	"github.com/jvanrhyn/bookshelf/internal/model"
)

// registerBookEndPoints registers the endpoints for the books API.
// It sets up the routes for interacting with the books data.
//
// Parameters:
//   - app: *fiber.App: The Fiber application instance.
//
// Returns:
//   - nil: If the registration is successful.
//   - error: If there is an error registering the endpoints.
func registerBookEndPoints(app *fiber.App) error {
	books := app.Group("books/")

	books.Get("/:id", getBookById)
	books.Get("/isbn/:isbn", getBookByISBN)
	books.Post("", createBook)
	books.Put("/:id", updateBook)
	books.Delete("/:id", deleteBook)

	return nil
}

// deleteBook deletes a book from the database based on the provided ID.
// It retrieves the ID from the context of the request, then uses the ID to delete the corresponding book from the database.
// If an error occurs during the deletion process, it logs the error and returns it.
// If the deletion is successful, it logs an info message indicating the deletion and returns nil.
//
// Parameters:
//   - ctx: *fiber.Ctx: The context of the request containing the ID of the book to be deleted.
//
// Returns:
//   - error: If there is an error deleting the book.
//   - nil: If the deletion is successful.
func deleteBook(ctx *fiber.Ctx) error {

	id, err := getIdFromContext(ctx)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = database.DeleteBook(id)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info("Deleted book", "id", id)
	return nil
}

// updateBook updates a book in the database based on the provided ID.
// It retrieves the ID from the context of the request, then uses the ID to update the corresponding book in the database.
// If an error occurs during the update process, it logs the error and returns it.
// If the update is successful, it logs an info message indicating the update and returns nil.
//
// Parameters:
//   - ctx: *fiber.Ctx: The context of the request containing the ID of the book to be updated.
//
// Returns:
//   - error: If there is an error updating the book.
//   - nil: If the update is successful.
func updateBook(ctx *fiber.Ctx) error {

	id, err := getIdFromContext(ctx)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	var updatedBook *model.Book
	err = ctx.BodyParser(&updatedBook)
	if err != nil {
		slog.Error("Failed to parse JSON body", "Error", err)
		return err
	}

	updatedBook.ID = uint(id)

	err = database.UpdateBook(updatedBook)
	if err != nil {
		slog.Error("Error in database update", "error", err)
		return err
	}

	return nil
}

// createBook is a function that is used as a handler for the POST "/" route.
// It attempts to create a new book in the database based on the JSON body received in the request.
// If the JSON body can be parsed into a Book object, the function responds with a success message in JSON format.
// If an error occurs during the parsing, the error is logged and returned.
// If an error occurs while responding with JSON, the error is logged.
// This function is expected to be used in a context where the router makes sure the request body exists.
// It returns an error that represents how the function execution went.
//
// The function takes a *fiber.Ctx object as its argument which represents the context of the request.
//
// Parameters:
//   - ctx: *fiber.Ctx: The context of the request containing the JSON body of the new book to be created.
//
// Returns:
//   - error: If there is an error creating the book.
//   - nil: If the creation is successful.
func createBook(ctx *fiber.Ctx) error {

	var book *model.Book
	err := ctx.BodyParser(&book)
	if err != nil {
		slog.Error("Failed to parse JSON body", "Error", err)
		return err
	}
	slog.Info("Received book from request", "Book", book)

	err = database.StoreBook(book)
	if err != nil {
		slog.Error("Error in database store", "error", err)
		return err
	}

	return nil
}

// getBookByISBN is a function that is used as a handler for the GET "/:isbn" route.
// It attempts to get a book's information based on its ISBN from the request parameters.
// If the ISBN is not empty and can be parsed to a string, the function responds with the ISBN in JSON format.
// If an error occurs during the parsing, the error is logged and returned.
// If an error occurs while responding with JSON, the error is logged.
// This function is expected to be used in a context where the router makes sure the ISBN parameter exists.
// It returns an error that represents how the function execution went.
//
// The function takes a *fiber.Ctx object as its argument which represents the context of the request.
//
// Parameters:
//   - ctx: *fiber.Ctx: The context of the request containing the ISBN of the book to be retrieved.
//
// Returns:
//   - error: If there is an error getting the book.
//   - nil: If the retrieval is successful.
func getBookByISBN(ctx *fiber.Ctx) error {

	isbn := ctx.Params("isbn")
	if isbn != "" {

		err, book := database.GetBookByISBN(isbn)
		if err != nil {
			slog.Error("Error in database get by ISBN", "error", err)
			return err
		}
		return ctx.JSON(book)
	}
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
		return err
	}

	err, book := database.GetBookById(id)
	if err != nil {
		slog.Error(err.Error())
	}

	err = ctx.JSON(book)
	if err != nil {
		slog.Error(err.Error(), "Could not marshal book into JSON", book)
		return err
	}
	return nil
}

// getIdFromContext retrieves the ID from the context of the request.
// It extracts the "id" parameter from the request's URL and attempts to parse it into an integer.
// If the "id" parameter is not empty, the function returns the parsed integer as an integer and nil.
// If an error occurs during the parsing, the function logs the error and returns the error and 0 as the integer.
// If the "id" parameter is empty, the function returns 0 and nil.
//
// Parameters:
//   - ctx: *fiber.Ctx: The context of the request containing the "id" parameter.
//
// Returns:
//   - (int, error): The parsed integer and nil if successful, or an error and 0 if an error occurs.
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
