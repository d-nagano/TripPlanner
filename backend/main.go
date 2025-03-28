package main

import (
	"trip-planner/infra"
	"trip-planner/requests"
	"trip-planner/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("error loading .env file")
	}

	appConfig, err := infra.InitAppConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	requests.InitValidator()

	e.Static("/api/tmp", "/var/www/backend/tmp")

	router.Router(e, appConfig)

	e.Logger.Fatal(e.Start(":8000"))
}
