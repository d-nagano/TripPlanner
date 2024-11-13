package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `grom:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}
