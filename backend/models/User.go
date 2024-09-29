package models

import "time"

type User struct {
	ID        int `grom:"primaryKey"`
	Name      string
	Email     string `grom:"unique"`
	Password  string
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}
