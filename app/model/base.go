package model

import (
	"time"
)

type BaseModel struct {
	ID        uint       `db:"primary_key"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
