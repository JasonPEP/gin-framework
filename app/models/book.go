package model

/*
Entity of book
*/
type Book struct {
	BaseModel
	Title     string  `gorm:"column:title"`
	AccountID int64   `gorm:"column:account_id"`
	Account   Account `gorm:"foreignkey:AccountID" `
}

func (Book) TableName() string {
	return "book"
}

func GetBook(id interface{}) (Book, error) {
	var book Book
	result := DB.Preload("Account").First(&book, id)
	return book, result.Error
}
