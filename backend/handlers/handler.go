package handlers

import (
	"trip-planner/infra"
)

type AppHandler struct {
	AppConfig *infra.AppConfig
}

func NewHandler(appConfig *infra.AppConfig) *AppHandler {
	return &AppHandler{
		AppConfig: appConfig,
	}
}
