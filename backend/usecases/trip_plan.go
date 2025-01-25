package usecases

import (
	"trip-planner/models"
	"trip-planner/repos"

	"gorm.io/gorm"
)

type TripPlanUseCase interface {
	GetTripPlans(userID string) (*models.TripPlanList, error)
}

type tripPlanUseCase struct {
	tripPlanRepo repos.TripPlanRepo
}

func NewTripPlanUseCase(db *gorm.DB) TripPlanUseCase {
	return &tripPlanUseCase{
		tripPlanRepo: repos.NewTripPlanRepo(db),
	}
}

func (tpu *tripPlanUseCase) GetTripPlans(userID string) (*models.TripPlanList, error) {
	tpl, err := tpu.tripPlanRepo.GetPlanByUserID(userID)
	if err != nil {
		return nil, err
	}

	return tpl, nil
}
