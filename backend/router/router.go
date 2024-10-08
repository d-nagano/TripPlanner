package router

import (
	"trip-planner/handlers"
	"trip-planner/infra"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, appConfig *infra.AppConfig) {
	handler := handlers.NewHandler(appConfig)

	e.GET("/api/prefectures", handler.GetPrefectures)
}
