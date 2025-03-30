package repos

import (
	"errors"
	"trip-planner/models"

	"gorm.io/gorm"
)

type TripPlanRepo interface {
	Create(tripPlan *models.TripPlan) error
	GetPlanByUserID(userID string) (*models.TripPlanList, error)
	GetPlanByID(tripID string) (*models.TripPlan, error)
}

type tripPlanRepo struct {
	db *gorm.DB
}

func NewTripPlanRepo(db *gorm.DB) TripPlanRepo {
	return &tripPlanRepo{db: db}
}

func (tpr *tripPlanRepo) Create(tripPlan *models.TripPlan) error {
	if tripPlan == nil {
		return errors.New("trip plan is nil")
	}

	if err := tpr.db.Create(tripPlan).Error; err != nil {
		return err
	}

	return nil
}

func (tpr *tripPlanRepo) GetPlanByUserID(userID string) (*models.TripPlanList, error) {
	if userID == "" {
		return nil, errors.New("user id is nil")
	}

	tripPlanList := &models.TripPlanList{}
	if err := tpr.db.Where("user_id = ?", userID).Find(&tripPlanList).Error; err != nil {
		return nil, err
	}

	return tripPlanList, nil
}

func (tpr *tripPlanRepo) GetPlanByID(tripID string) (*models.TripPlan, error) {
	if tripID == "" {
		return nil, errors.New("trip plan id is nil")
	}

	tripPlan := &models.TripPlan{}
	if err := tpr.db.Where("id = ?", tripID).First(&tripPlan).Error; err != nil {
		return nil, err
	}

	return tripPlan, nil
}
