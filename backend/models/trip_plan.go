package models

import (
	"time"

	"gorm.io/gorm"
)

type TripPlan struct {
	ID            string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	Title         string    `gorm:"type:varchar(64);not null" json:"title"`
	Destination   string    `gorm:"type:varchar(64);not null" json:"destination"`
	DepartureDate time.Time `gorm:"type:datetime" json:"departure_date"`
	ArrivalDate   time.Time `gorm:"type:datetime" json:"arrival_date"`
	UserID        string    `gorm:"type:varchar(36);not null" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type TripPlanList []TripPlan
