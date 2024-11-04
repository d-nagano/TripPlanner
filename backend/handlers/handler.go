package handlers

import (
	"trip-planner/infra"

	"github.com/rs/zerolog"
)

type AppHandler struct {
	AppConfig *infra.AppConfig
	Logger    zerolog.Logger
}

func NewHandler(appConfig *infra.AppConfig, logger zerolog.Logger) *AppHandler {
	return &AppHandler{
		AppConfig: appConfig,
		Logger:    logger,
	}
}
