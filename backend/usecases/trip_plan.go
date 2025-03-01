package usecases

import (
	"trip-planner/models"
	"trip-planner/repos"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TripPlanUseCase interface {
	GetTripPlans(userID string) (*models.TripPlanList, error)
	RegisterTripPlan(tripPlan *models.TripPlan) (string, error)
}

type tripPlanUseCase struct {
	db           *gorm.DB
	tripPlanRepo repos.TripPlanRepo
	tripDayRepo  repos.TripDayRepo
}

func NewTripPlanUseCase(db *gorm.DB) TripPlanUseCase {
	return &tripPlanUseCase{
		db:           db,
		tripPlanRepo: repos.NewTripPlanRepo(db),
		tripDayRepo:  repos.NewTripDayRepoRepo(db),
	}
}

func (tpu *tripPlanUseCase) GetTripPlans(userID string) (*models.TripPlanList, error) {
	tpl, err := tpu.tripPlanRepo.GetPlanByUserID(userID)
	if err != nil {
		return nil, err
	}

	return tpl, nil
}

func (tpu *tripPlanUseCase) RegisterTripPlan(tripPlan *models.TripPlan) (string, error) {
	tripID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	tripPlan.ID = tripID.String()
	startDate := tripPlan.StartDate
	diffDays := int(tripPlan.EndDate.Sub(startDate).Hours() / 24)

	tripDayList := make(models.TripDayList, 0, diffDays+1)
	for i := 0; i <= diffDays; i++ {
		tripDayID, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}

		tripDay := &models.TripDay{
			ID:         tripDayID.String(),
			DayNumber:  i + 1,
			PlanDate:   startDate.AddDate(0, 0, i),
			TripPlanID: tripPlan.ID,
		}
		tripDayList = append(tripDayList, tripDay)
	}

	if err = tpu.db.Transaction(func(tx *gorm.DB) error {
		txTripPlanRepo := repos.NewTripPlanRepo(tx)
		txTripDayRepo := repos.NewTripDayRepoRepo(tx)

		if err := txTripPlanRepo.Create(tripPlan); err != nil {
			return err
		}
		if err := txTripDayRepo.Creates(tripDayList); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return "", err
	}

	return tripPlan.ID, nil
}
