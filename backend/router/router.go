package router

import (
	"trip-planner/handlers"
	"trip-planner/infra"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo, appConfig *infra.AppConfig) {
	handler := handlers.NewHandler(appConfig)

	e.Use(middleware.CORS())
	e.GET("/api/prefectures", handler.GetPrefectures)
}
