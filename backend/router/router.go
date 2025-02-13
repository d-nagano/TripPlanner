package router

import (
	"log/slog"
	"net/http"
	"os"
	"trip-planner/common"
	"trip-planner/handlers"
	"trip-planner/infra"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"golang.org/x/time/rate"
)

func Router(e *echo.Echo, appConfig *infra.AppConfig) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logConfig := slogecho.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
	}
	e.Use(slogecho.NewWithConfig(logger, logConfig))

	handler := handlers.NewHandler(appConfig)
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
			common.LogError(c, err)
			return c.JSON(http.StatusUnauthorized, "認証エラー")
		},
	}
	g.Use(echojwt.WithConfig(config))
	g.GET("", handler.GetTripPlans)
}
