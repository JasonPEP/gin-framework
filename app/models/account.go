package model

/*
Entity of account
*/
type Account struct {
	BaseModel
	Name string `gorm:"size:50;column:name"`
}

func (Account) TableName() string {
	return "account"
}

func GetAccount(id interface{}) (Account, error) {
	var account Account
	result := DB.First(&account, id)
	return account, result.Error
}
