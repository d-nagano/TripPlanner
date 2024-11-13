package requests

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Name     string `validate:"required,max=64"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,password,min=8,max=64"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

var (
	ErrEmptyName       = errors.New("name is empty")
	ErrLongName        = errors.New("name is too long")
	ErrEmptyEmail      = errors.New("email is empty")
	ErrInvalidEmail    = errors.New("email is invalid")
	ErrEmptyPassword   = errors.New("password is empty")
	ErrInvalidPassword = errors.New("password is invalid")
	ErrShortPassword   = errors.New("password is too short")
	ErrLongPassword    = errors.New("password is too long")
)

func (sr *SignUpRequest) Validate() []error {
	var validationErrors []error
	if err := sr.validateName(); err != nil {
		validationErrors = append(validationErrors, err)
	}
	if err := sr.validateEmail(); err != nil {
		validationErrors = append(validationErrors, err)
	}
	if err := sr.validatePassword(); err != nil {
		validationErrors = append(validationErrors, err)
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	// todo: パスワードの正規表現追加

	return nil
}

func (sr *SignUpRequest) validateName() error {
	sr.Name = strings.TrimSpace(sr.Name)

	if err := Validate.Var(sr.Name, "required,max=64"); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return ErrEmptyName
			case "max":
				return ErrLongName
			}
		}
	}

	return nil
}

func (sr *SignUpRequest) validateEmail() error {
	sr.Email = strings.TrimSpace(sr.Email)

	if err := Validate.Var(sr.Email, "required,email"); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return ErrEmptyEmail
			case "email":
				return ErrInvalidEmail
			}
		}
	}

	return nil
}

func (sr *SignUpRequest) validatePassword() error {
	sr.Password = strings.TrimSpace(sr.Password)

	Validate.RegisterValidation("password", checkPassword)
	if err := Validate.Var(sr.Password, "required,password,min=8,max=64"); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return ErrEmptyPassword
			case "password":
				return ErrInvalidPassword
			case "min":
				return ErrShortPassword
			case "max":
				return ErrLongPassword
			}
		}
	}

	return nil
}

func checkPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	isLowerCase := checkRegexp("[a-z]", password)
	isUpperCase := checkRegexp("[A-Z]", password)
	isNumeric := checkRegexp("[0-9]", password)
	availableChar := checkRegexp(`^[0-9a-zA-Z\-^$*.@]+$`, password)
	isSymbol := checkRegexp(`[\-^$*.@]`, password)

	return isLowerCase && isUpperCase && isNumeric && availableChar && isSymbol
}

func checkRegexp(reg, str string) bool {
	r := regexp.MustCompile(reg).Match([]byte(str))
	return r
}
