package repos

import (
	"errors"
	"trip-planner/models"

	"gorm.io/gorm"
)

type TripDayRepo interface {
	Creates(tripDayList models.TripDayList) error
}

type tripDayRepo struct {
	db *gorm.DB
}

func NewTripDayRepoRepo(db *gorm.DB) TripDayRepo {
	return &tripDayRepo{db: db}
}

func (tdr *tripDayRepo) Creates(tripDayList models.TripDayList) error {
	if tripDayList == nil {
		return errors.New("trip day list is nil")
	}

	if err := tdr.db.Create(tripDayList).Error; err != nil {
		return err
	}

	return nil
}
