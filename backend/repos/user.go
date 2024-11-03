package repos

import (
	"errors"
	"trip-planner/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(*models.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}
