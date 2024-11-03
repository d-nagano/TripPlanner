package handlers

import (
	"net/http"
	"strings"
	"trip-planner/models"
	"trip-planner/requests"
	"trip-planner/usecases"

	"github.com/labstack/echo/v4"
)

func (h *AppHandler) SignUp(c echo.Context) error {
	var req requests.SignUpRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "リクエストの形式が正しくありません。")
	}
	if errs := req.Validate(); errs != nil {
		var errorMessages []string
		for _, err := range errs {
			c.Logger().Warn(err.Error())
			switch err {
			case requests.ErrEmptyName:
				errorMessages = append(errorMessages, "名前は必須です。")
			case requests.ErrLongName:
				errorMessages = append(errorMessages, "名前が長すぎます。")
			case requests.ErrEmptyEmail:
				errorMessages = append(errorMessages, "メールアドレスは必須です。")
			case requests.ErrInvalidEmail:
				errorMessages = append(errorMessages, "メールアドレスの形式が正しくありません。")
			case requests.ErrEmptyPassword:
				errorMessages = append(errorMessages, "パスワードは必須です。")
			case requests.ErrInvalidPassword:
				errorMessages = append(errorMessages, "パスワードは大文字、小文字、数字、特殊文字を含める必要があります。")
			case requests.ErrShortPassword:
				errorMessages = append(errorMessages, "パスワードが短すぎます。")
			case requests.ErrLongPassword:
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
		switch err {
		case usecases.ErrDuplicateEmail:
			c.Logger().Warn(err.Error())
			return c.JSON(http.StatusBadRequest, "このメールアドレスは既に登録済みです。")
		default:
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, "ユーザーの登録中にエラーが発生しました。")
		}

	}

	return c.NoContent(http.StatusNoContent)
}
