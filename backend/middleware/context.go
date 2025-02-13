package middleware

import (
	"net/http"
	"trip-planner/common"
	"trip-planner/infra"
	"trip-planner/repos"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ContextMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*infra.JwtCustomClaims)
			userID := claims.UserID

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
