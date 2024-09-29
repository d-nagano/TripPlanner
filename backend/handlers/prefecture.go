package handlers

import (
	"net/http"
	"trip-planner/usecases"

	"github.com/labstack/echo/v4"
)

func (h *AppHandler) GetPrefectures(c echo.Context) error {
	u := usecases.NewPrefectureUseCase(h.AppConfig.DB)

	resp, err := u.GetPrefectures()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}
