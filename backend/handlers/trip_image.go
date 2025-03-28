package handlers

import (
	"net/http"
	"trip-planner/common"
	"trip-planner/infra"
	"trip-planner/usecases"

	"github.com/labstack/echo/v4"
)

func (h *AppHandler) UploadImage(c echo.Context) error {
	ctx := c.(*infra.CustomContext)
	
	tripID := c.Param("tripID")
	file, err := c.FormFile("file")
	if err != nil {
		common.LogError(c, err)
		return c.JSON(http.StatusInternalServerError, "ファイルの取得中にエラーが発生しました。")
	}

	tiu := usecases.NewTripImageUseCase(h.AppConfig.DB)
	if err := tiu.UploadFile(ctx.ActingUser.ID, tripID, file); err != nil {
		common.LogError(c, err.Error)
		c.JSON(err.StatusCode, err.Message)
	}

	return c.NoContent(http.StatusNoContent)
}
