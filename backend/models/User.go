package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	Name      string `gorm:"type:varchar(64)" json:"name"`
	Email     string `gorm:"type:varchar(256);unique" json:"email"`
	Password  string `gorm:"type:varchar(64)" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
