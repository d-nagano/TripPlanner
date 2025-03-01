package usecases

import (
	"encoding/json"
	"fmt"
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
	tripPlanRepo repos.TripPlanRepo
	tripDayRepo  repos.TripDayRepo
}

func NewTripPlanUseCase(db *gorm.DB) TripPlanUseCase {
	return &tripPlanUseCase{
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
	diffDays := tripPlan.EndDate.Sub(startDate).Hours() / 24

	tripDayList := make(models.TripDayList, 0, int(diffDays)+1)
	for i := 0; i <= int(diffDays); i++ {
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

	out, _ := json.Marshal(tripDayList)
	fmt.Println(string(out))

	if err := tpu.tripPlanRepo.Create(tripPlan); err != nil {
		return "", err
	}
	if err := tpu.tripDayRepo.Creates(tripDayList); err != nil {
		return "", err
	}

	return tripPlan.ID, nil
}
