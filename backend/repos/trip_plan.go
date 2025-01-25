package repos

import (
	"fmt"
	"trip-planner/models"

	"gorm.io/gorm"
)

type TripPlanRepo interface {
	GetPlanByUserID(userID string) (*models.TripPlanList, error)
}

type tripPlanRepo struct {
	db *gorm.DB
}

func NewTripPlanRepo(db *gorm.DB) TripPlanRepo {
	return &tripPlanRepo{db: db}
}

func (tpr *tripPlanRepo) GetPlanByUserID(userID string) (*models.TripPlanList, error) {
	var tripPlanList *models.TripPlanList
	if err := tpr.db.Where("user_id = ?", userID).Find(&tripPlanList).Error; err != nil {
		return nil, err
	}
	fmt.Println(tripPlanList)

	return tripPlanList, nil
}
