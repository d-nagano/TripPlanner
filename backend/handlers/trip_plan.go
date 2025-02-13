package handlers

import (
	"net/http"
	"trip-planner/common"
	"trip-planner/infra"
	"trip-planner/usecases"

	"github.com/labstack/echo/v4"
)

func (h *AppHandler) GetTripPlans(c echo.Context) error {
	ctx := c.(*infra.CustomContext)

	tpu := usecases.NewTripPlanUseCase(h.AppConfig.DB)
	tpl, err := tpu.GetTripPlans(ctx.ActingUser.ID)
	if err != nil {
		common.LogError(c, err)
		return c.JSON(http.StatusInternalServerError, "データ取得中にエラーが発生しました。")
	}

	return c.JSON(http.StatusOK, tpl)
}
