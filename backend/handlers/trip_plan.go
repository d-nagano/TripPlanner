package handlers

import (
	"net/http"
	"trip-planner/common"
	"trip-planner/infra"
	"trip-planner/usecases"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *AppHandler) GetTripPlans(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*infra.JwtCustomClaims)
	userID := claims.UserID

	tpu := usecases.NewTripPlanUseCase(h.AppConfig.DB)
	tpl, err := tpu.GetTripPlans(userID)
	if err != nil {
		common.LogError(c, err)
		return c.JSON(http.StatusInternalServerError, "データ取得中にエラーが発生しました。")
	}

	return c.JSON(http.StatusOK, tpl)
}
