package middleware

import (
	"log/slog"
	"net/http"
	"trip-planner/common"
	"trip-planner/infra"
	"trip-planner/repos"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
	"gorm.io/gorm"
)

func ContextMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := validateToken(c)
			if  userID == ""{
				slogecho.AddCustomAttributes(c, slog.String("error msg", "invalid token"))
				return c.JSON(http.StatusUnauthorized, "トークンが不正です。")
			}

			uRepo := repos.NewUserRepo(db)
			actingUser, err := uRepo.FindByID(userID)
			if err != nil {
				common.LogError(c, err)
				return c.JSON(http.StatusNotFound, "ユーザーの取得中にエラーが発生しました")
			}
			ctx := &infra.CustomContext{
				Context:    c,
				ActingUser: actingUser,
			}
			c.Set("acting_user", actingUser)
			return next(ctx)
		}
	}
}

func validateToken(c echo.Context) string {
	userToken := c.Get("user")
	if userToken == nil {
		return ""
	}

	user, ok := userToken.(*jwt.Token)
	if !ok {
		return ""
	}

	claims, ok := user.Claims.(*infra.JwtCustomClaims)
	if !ok {
		return ""
	}

	return claims.UserID
}
