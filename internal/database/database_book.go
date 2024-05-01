package database

import (
	"github.com/go-errors/errors"
	"github.com/jvanrhyn/bookshelf/internal/model"
	"log/slog"
)

func StoreBook(book *model.Book) error {

	tx := db.Create(&book)
	if tx.Error != nil {
		stack := err.(*errors.Error).ErrorStack()
		slog.Error("error while recording lookup", "error", err, "stacktrace", stack)
	}
	return nil
}

func DeleteBook(id int) error {

	book := model.Book{}

	result := db.First(&book, id)
	if result.Error != nil {
		slog.Error("Error retrieving the book", "error", result.Error)
		return err
	}

	tx := db.Delete(&book)
	if tx.Error != nil {
		stack := err.(*errors.Error).ErrorStack()
		slog.Error("error while recording lookup", "error", err, "stacktrace", stack)
	}

	slog.Info("Delete book", "id", id)
	return nil
}

func GetBookById(id int) (error, *model.Book) {
	book := model.Book{}

	result := db.First(&book, id)
	if result.Error != nil {
		slog.Error("Error retrieving the book", "error", result.Error)
		return err, &model.Book{}
	}

	return nil, &book
}

// UpdateBook updates an existing book in the database.
//
// Parameters:
//   - book: A pointer to the model.Book struct that represents the book to be updated.
//
// Returns:
//   - error: An error if the update fails.
//
// Example:
//
//	book := &model.Book{ID: 1, Title: "New Title"}
//	err := database.UpdateBook(book)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Book updated successfully.")
func UpdateBook(book *model.Book) error {

	err := db.Save(book).Error
	if err != nil {
		slog.Error("Error in database update", "error", err)
		return err
	}

	return nil
}

// GetBookByISBN retrieves a book from the database based on its ISBN.
//
// Parameters:
//   - isbn: A string that represents the ISBN of the book to be retrieved.
//
// Returns:
//   - error: An error if the retrieval fails.
//   - *model.Book: A pointer to the model.Book struct that represents the book retrieved from the database.
func GetBookByISBN(isbn string) (error, *model.Book) {

	book := model.Book{}

	result := db.First(&book, "isbn = ?", isbn)
	if result.Error != nil {
		slog.Error("Error retrieving the book", "error", result.Error)
		return err, &model.Book{}
	}

	return nil, &book
}
