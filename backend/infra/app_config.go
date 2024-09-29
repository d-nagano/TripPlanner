package infra

import (
	"errors"

	"gorm.io/gorm"
)

type AppConfig struct {
	DB *gorm.DB
}

func InitAppConfig() (*AppConfig, error) {
	db, err := initDB()
	if err != nil {
		return nil, errors.New("conect db failed")
	}

	return &AppConfig{
		DB: db,
	}, nil
}
