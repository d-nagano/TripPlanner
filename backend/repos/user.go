package repos

import (
	"errors"
	"trip-planner/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(*models.User) error
	FindByEmail(string) (*models.User, error)
	FindByID(string) (*models.User, error)
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

func (ur *userRepo) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) FindByID(id string) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
