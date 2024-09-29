package repos

import (
	"trip-planner/models"

	"gorm.io/gorm"
)

type PrefectureRepo interface {
	GetAll() (models.Prefectures, error)
}

type prefectureRepo struct {
	db *gorm.DB
}

func NewPrefectureRepo(db *gorm.DB) PrefectureRepo {
	return &prefectureRepo{db: db}
}

func (pr *prefectureRepo) GetAll() (models.Prefectures, error) {
	var prefectures models.Prefectures
	if err := pr.db.Order("id").Find(&prefectures).Error; err != nil {
		return nil, err
	}

	return prefectures, nil
}
