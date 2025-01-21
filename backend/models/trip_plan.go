package models

import (
	"time"

	"gorm.io/gorm"
)

type TripPlan struct {
	gorm.Model
	ID            int       `gorm:"AUTO_INCREMENT" json:"id"`
	Title         string    `json:"title"`
	Destination   string    `json:"destination"`
	DepartureDate time.Time `gorm:"type:datetime" json:"departure_date"`
	ArrivalDate   time.Time `gorm:"type:datetime" json:"arrival_date"`
	UserID        string    `json:"user_id"`
	User          User      `gorm:"foreignKey:UserID"`
}

type TripPlanList []TripPlan
