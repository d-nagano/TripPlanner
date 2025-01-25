package router

import (
	"net/http"
	"os"
	"trip-planner/handlers"
	"trip-planner/infra"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

func Router(e *echo.Echo, appConfig *infra.AppConfig) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))

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
	e.POST("/api/logout", handler.Logout)

	g := e.Group("/api/trip-plan")

	config := echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SECRET_KEY")),
		TokenLookup: "cookie:token",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(infra.JwtCustomClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			logger.Warn().Msg("invalid token")
			return c.JSON(http.StatusUnauthorized, "認証エラー")
		},
	}
	g.Use(echojwt.WithConfig(config))
	g.GET("", handler.GetTripPlans)
}
