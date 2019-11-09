package model

import (
	"time"
)

/*
Base entity
*/
type ID uint
type BaseModel struct {
	ID          uint
	CreatedTime time.Time
	UpdatedTime time.Time
	DeletedTime *time.Time
}

func (b *BaseModel) FindOne() int {
	return 1
}
