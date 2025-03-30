package common

import "net/http"

type AppError struct {
	StatusCode int
	Message    string
	Error      error
}

func NotFoundError(err error) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    "データが見つかりません。",
		Error:      err,
	}
}

func BadRequestError(err error) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    "リクエストが不正です。",
		Error:      err,
	}
}

func InternalServerError(err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    "サーバー内部でエラーが発生しました。",
		Error:      err,
	}
}
