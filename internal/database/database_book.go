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
