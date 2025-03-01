package common

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
)

func LogError(c echo.Context, err error) {
	slogecho.AddCustomAttributes(c, slog.String("error msg", err.Error()))
}

func LogErrors(c echo.Context, errs []error) {
	for _, err := range errs {
		slogecho.AddCustomAttributes(c, slog.String("error msg", err.Error()))
	}
}
