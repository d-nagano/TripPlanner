package common

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
)

func LogError(c echo.Context, err error) {
	slogecho.AddCustomAttributes(c, slog.String("error msg", err.Error()))
}
