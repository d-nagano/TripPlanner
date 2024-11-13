package router

import (
	"os"
	"trip-planner/handlers"
	"trip-planner/infra"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func Router(e *echo.Echo, appConfig *infra.AppConfig) {
	e.Use(middleware.CORS())

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	handler := handlers.NewHandler(appConfig, logger)
	e.GET("/api/prefectures", handler.GetPrefectures)
	e.POST("/api/signup", handler.SignUp)
	e.POST("/api/login", handler.Login)

	g := e.Group("/trip-planner")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(infra.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	g.Use(echojwt.WithConfig(config))
}
