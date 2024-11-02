package repos

import (
	"errors"
	"trip-planner/models"

	"github.com/go-sql-driver/mysql"
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
	if err := ur.db.Create(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				return errors.New("dupulicate entry")
			default:
				// nope
			}
		} else {
			return err
		}
	}

	return nil
}
