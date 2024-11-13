package usecases

import (
	"errors"
	"time"
	"trip-planner/infra"
	"trip-planner/models"
	"trip-planner/repos"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	SignUp(user models.User) error
	Login(user models.User) (string, error)
}

type userUseCase struct {
	UserRepo repos.UserRepo
}

func NewUserUseCase(db *gorm.DB) UserUseCase {
	return &userUseCase{
		UserRepo: repos.NewUserRepo(db),
	}
}

var (
	ErrDuplicateEmail   = errors.New("duplicate entry")
	ErrPasswordMismatch = errors.New("password is mismatch")
)

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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

func (uu *userUseCase) Login(user models.User) (string, error) {
	registeredUser, err := uu.UserRepo.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if registeredUser == nil {
		return "", ErrPasswordMismatch
	}

	err = bcrypt.CompareHashAndPassword([]byte(registeredUser.Password), []byte(user.Password))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return "", ErrPasswordMismatch
		default:
			return "", err
		}
	}

	claims := &infra.JwtCustomClaims{
		UserID: registeredUser.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
