package model

/*
Entity of account
*/
type Account struct {
	BaseModel
	Name string `db:"name"`
}
