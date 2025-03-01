package handlers

import (
	"log/slog"
	"net/http"
	"trip-planner/infra"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
)

func (h *AppHandler) UploadImage(c echo.Context) error {
	tripID := c.Param("tripID")

	ctx := c.(*infra.CustomContext)
	slogecho.AddCustomAttributes(c, slog.String("user_id", ctx.ActingUser.ID))
	slogecho.AddCustomAttributes(c, slog.String("trip_plan_id", tripID))

	// ToDo: ファイルのアップロード機能追加

	return c.NoContent(http.StatusNoContent)
}
