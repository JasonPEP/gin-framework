package model

import (
	"time"
)

/*
Base entity
*/
type BaseModel struct {
	ID          int64      `gorm:"primary_key,AUTO_INCREMENT;column:id"`
	CreatedTime time.Time  `gorm:"column:created_time;default:'NOW()'"`
	UpdatedTime time.Time  `gorm:"column:updated_time;default:'NOW()'"`
	DeletedTime *time.Time `gorm:"column:deleted_time;default:'null'" sql:"index"`
}
