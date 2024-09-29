package models

import "time"

type TripPlan struct {
	ID            int `gorm:"AUTO_INCREMENT"`
	Title         string
	Destination   string
	DepartureDate time.Time `gorm:"type:datetime(6)"`
	ArrivalDate   time.Time `gorm:"type:datetime(6)"`
	CreatedAt     time.Time `gorm:"type:datetime(6)"`
	UpdatedAt     time.Time `gorm:"type:datetime(6)"`
}
