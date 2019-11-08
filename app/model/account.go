package model

import (
	"database/sql"
)

type Account struct {
	BaseModel
	Name sql.NullString `db:"name"`
}
