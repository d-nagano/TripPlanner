package infra

import (
	"trip-planner/models"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	ActingUser *models.User
}
