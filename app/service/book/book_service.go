package bookService

import (
	"errors"
	model "gin-framework/app/models"
)

// Create a book
func NewBook(title string, accountId int64) (book *model.Book, err error) {
	newBook := &model.Book{
		Title:     title,
		AccountID: accountId,
	}
	can := model.DB.NewRecord(newBook)

	if can {
		model.DB.Create(&newBook)
	} else {
		return nil, errors.New("create a book fail")
	}
	return newBook, nil
}
