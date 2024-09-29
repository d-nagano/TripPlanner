package usecases

import (
	"trip-planner/models"
	"trip-planner/repos"

	"gorm.io/gorm"
)

type PrefectureUseCase interface {
	GetPrefectures() (models.Prefectures, error)
}

type prefectureUseCase struct {
	prefectureRepo repos.PrefectureRepo
}

func NewPrefectureUseCase(db *gorm.DB) PrefectureUseCase {
	return &prefectureUseCase{
		prefectureRepo: repos.NewPrefectureRepo(db),
	}
}

func (pu *prefectureUseCase) GetPrefectures() (models.Prefectures, error) {
	prefectures, err := pu.prefectureRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return prefectures, nil
}
