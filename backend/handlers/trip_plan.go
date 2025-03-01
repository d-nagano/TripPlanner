package handlers

import (
	"net/http"
	"strings"
	"trip-planner/common"
	"trip-planner/infra"
	"trip-planner/models"
	"trip-planner/requests"
	"trip-planner/responses"
	"trip-planner/usecases"

	"github.com/labstack/echo/v4"
)

func (h *AppHandler) GetTripPlans(c echo.Context) error {
	ctx := c.(*infra.CustomContext)

	tpu := usecases.NewTripPlanUseCase(h.AppConfig.DB)
	tpl, err := tpu.GetTripPlans(ctx.ActingUser.ID)
	if err != nil {
		common.LogError(c, err)
		return c.JSON(http.StatusInternalServerError, "データ取得中にエラーが発生しました。")
	}

	return c.JSON(http.StatusOK, tpl)
}

func (h *AppHandler) RegisterTripPlan(c echo.Context) error {
	ctx := c.(*infra.CustomContext)
	var req *requests.TripPlanRequest
	if err := c.Bind(&req); err != nil {
		common.LogError(c, err)
		return c.JSON(http.StatusBadRequest, "リクエストの形式が正しくありません。")
	}

	startDate, endDate, errs := req.Validate()
	if errs != nil {
		var errorMessages []string
		for _, err := range errs {
			switch err {
			case requests.ErrEmptyTitle:
				errorMessages = append(errorMessages, "タイトルは必須です。")
			case requests.ErrLongTitle:
				errorMessages = append(errorMessages, "タイトルが長すぎます。")
			case requests.ErrEmptyDestination:
				errorMessages = append(errorMessages, "目的地は必須です。")
			case requests.ErrLongDestination:
				errorMessages = append(errorMessages, "目的地の名称が長すぎます。")
			case requests.ErrInvalidStartDate:
				errorMessages = append(errorMessages, "出発日の日付が不正です。")
			case requests.ErrInvalidEndDate:
				errorMessages = append(errorMessages, "到着日の日付が不正です。")
			case requests.ErrInvalidOrder:
				errorMessages = append(errorMessages, "出発日は到着日よりも前の必要があります。")
			default:
				errorMessages = append(errorMessages, "入力エラーが発生しました。")
			}
		}

		common.LogErrors(c, errs)
		return c.JSON(http.StatusBadRequest, strings.Join(errorMessages, "\n"))
	}

	tp := &models.TripPlan{
		Title:       req.Title,
		Destination: req.Destination,
		StartDate:   startDate,
		EndDate:     endDate,
		UserID:      ctx.ActingUser.ID,
	}

	tpu := usecases.NewTripPlanUseCase(h.AppConfig.DB)
	tripID, err := tpu.RegisterTripPlan(tp)
	if err != nil {
		common.LogError(c, err)
		return c.JSON(http.StatusInternalServerError, "データ登録中にエラーが発生しました。")
	}

	return c.JSON(http.StatusOK, responses.RegisterTripPlanResponse{
		ID: tripID,
	})
}
