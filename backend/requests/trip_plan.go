package requests

import (
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type TripPlanRequest struct {
	Title       string `json:"title" validate:"required,max=64"`
	Destination string `json:"destination" validate:"required,max=64"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

var (
	ErrEmptyTitle       = errors.New("title is empty")
	ErrLongTitle        = errors.New("title is too long")
	ErrEmptyDestination = errors.New("destination is empty")
	ErrLongDestination  = errors.New("destination is too long")
	ErrInvalidStartDate = errors.New("departure date is invalid format")
	ErrInvalidEndDate   = errors.New("arrival date is invalid format")
	ErrInvalidOrder     = errors.New("departure date must be before arrival date")
)

const layout = "2006-01-02"

func (tpr *TripPlanRequest) Validate() (time.Time, time.Time, []error) {
	var validationErrors []error
	if err := tpr.validateTitle(); err != nil {
		validationErrors = append(validationErrors, err)
	}
	if err := tpr.validateDestination(); err != nil {
		validationErrors = append(validationErrors, err)
	}
	startDate, endDate, errs := tpr.validateDates()
	if len(errs) > 0 {
		validationErrors = append(validationErrors, errs...)
	}

	if len(validationErrors) > 0 {
		return time.Time{}, time.Time{}, validationErrors
	}

	return startDate, endDate, nil
}

func (tpr *TripPlanRequest) validateTitle() error {
	tpr.Title = strings.TrimSpace(tpr.Title)

	if err := Validate.Var(tpr.Title, "required,max=64"); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return ErrEmptyTitle
			case "max":
				return ErrLongTitle
			}
		}
	}

	return nil
}

func (tpr *TripPlanRequest) validateDestination() error {
	tpr.Destination = strings.TrimSpace(tpr.Destination)

	if err := Validate.Var(tpr.Destination, "required,max=64"); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return ErrEmptyDestination
			case "max":
				return ErrLongDestination
			}
		}
	}

	return nil
}

func (tpr *TripPlanRequest) validateDates() (time.Time, time.Time, []error) {
	var validationDateErrors []error

	dep, err := time.Parse(layout, tpr.StartDate)
	if err != nil {
		validationDateErrors = append(validationDateErrors, ErrInvalidStartDate)
	}
	arr, err := time.Parse(layout, tpr.EndDate)
	if err != nil {
		validationDateErrors = append(validationDateErrors, ErrInvalidEndDate)
	}

	if dep.After(arr) {
		validationDateErrors = append(validationDateErrors, ErrInvalidOrder)
	}

	return dep, arr, validationDateErrors
}
