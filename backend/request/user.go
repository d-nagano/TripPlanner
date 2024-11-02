package request

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Name     string `validate:"required,max=64"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=64"`
}

var validate = validator.New()

var (
	ErrEmptyName     = errors.New("name is empty")
	ErrLongName      = errors.New("name is too long")
	ErrEmptyEmail    = errors.New("email is empty")
	ErrInvalidEmail  = errors.New("email is invalid")
	ErrEmptyPassword = errors.New("password is empty")
	ErrShortPassword = errors.New("password is too short")
	ErrLongPassword  = errors.New("password is too long")
)

func (sr *SignUpRequest) Validate() []error {
	sr.Name = strings.TrimSpace(sr.Name)
	sr.Email = strings.TrimSpace(sr.Email)
	sr.Password = strings.TrimSpace(sr.Password)

	var validationErrors []error
	if err := validate.Struct(sr); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Name":
				switch err.Tag() {
				case "required":
					validationErrors = append(validationErrors, ErrEmptyName)
				case "max":
					validationErrors = append(validationErrors, ErrLongName)
				}
			case "Email":
				switch err.Tag() {
				case "required":
					validationErrors = append(validationErrors, ErrEmptyEmail)
				case "email":
					validationErrors = append(validationErrors, ErrInvalidEmail)
				}
			case "Password":
				switch err.Tag() {
				case "required":
					validationErrors = append(validationErrors, ErrEmptyPassword)
				case "min":
					validationErrors = append(validationErrors, ErrShortPassword)
				case "max":
					validationErrors = append(validationErrors, ErrLongPassword)
				}
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	// todo: パスワードの正規表現追加

	return nil
}
