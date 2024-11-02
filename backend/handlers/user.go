package handlers

import (
	"net/http"
	"strings"
	"trip-planner/models"
	"trip-planner/request"
	"trip-planner/usecases"

	"github.com/labstack/echo/v4"
)

func (h *AppHandler) SignUp(c echo.Context) error {
	var req request.SignUpRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "リクエストの形式が正しくありません。")
	}
	if errs := req.Validate(); errs != nil {
		var errorMessages []string
		for _, err := range errs {
			c.Logger().Warn(err.Error())
			switch err {
			case request.ErrEmptyName:
				errorMessages = append(errorMessages, "名前は必須です。")
			case request.ErrLongName:
				errorMessages = append(errorMessages, "名前が長すぎます。")
			case request.ErrEmptyEmail:
				errorMessages = append(errorMessages, "メールアドレスは必須です。")
			case request.ErrInvalidEmail:
				errorMessages = append(errorMessages, "メールアドレスの形式が正しくありません。")
			case request.ErrEmptyPassword:
				errorMessages = append(errorMessages, "パスワードは必須です。")
			case request.ErrShortPassword:
				errorMessages = append(errorMessages, "パスワードが短すぎます。")
			case request.ErrLongPassword:
				errorMessages = append(errorMessages, "パスワードが長すぎます。")
			default:
				errorMessages = append(errorMessages, "入力エラーが発生しました。")
			}
		}

		return c.JSON(http.StatusBadRequest, strings.Join(errorMessages, "\n"))
	}

	uu := usecases.NewUserUseCase(h.AppConfig.DB)

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := uu.SignUp(user); err != nil {
		switch err.Error() {
		case "dupulicate entry":
			c.Logger().Warn(err.Error())
			return c.JSON(http.StatusBadRequest, "このメールアドレスは既に登録済みです。")
		default:
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "ユーザーの登録中にエラーが発生しました。")
		}

	}

	return c.NoContent(http.StatusNoContent)
}
