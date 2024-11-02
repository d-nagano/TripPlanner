package usecases

import (
	"trip-planner/models"
	"trip-planner/repos"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	SignUp(user models.User) error
}

type userUseCase struct {
	UserRepo repos.UserRepo
}

func NewUserUseCase(db *gorm.DB) UserUseCase {
	return &userUseCase{
		UserRepo: repos.NewUserRepo(db),
	}
}

func (uu *userUseCase) SignUp(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.UserRepo.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}
