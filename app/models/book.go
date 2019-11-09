package model

/*
Entity of book
*/
type Book struct {
	BaseModel
	Account Account `db:"account_id"`
}
