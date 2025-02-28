package models

import (
	"time"
)

type TripDay struct {
	ID         string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	DayNumber  int       `gorm:"type:int" json:"day_number"`
	PlanDate   time.Time `gorm:"type:date" json:"plan_date"`
	TripPlanID string    `gorm:"type:varchar(36);not null" json:"trip_plan_id"`
	TripPlan   TripPlan  `gorm:"foreignKey:TripPlanID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
