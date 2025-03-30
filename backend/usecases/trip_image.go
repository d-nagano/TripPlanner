package usecases

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"trip-planner/common"
	"trip-planner/repos"

	"gorm.io/gorm"
)

type TripImageUseCase interface {
	UploadFile(userID, tripID string, file *multipart.FileHeader) *common.AppError
}

type tripImageUseCase struct {
	db           *gorm.DB
	tripPlanRepo repos.TripPlanRepo
}

func NewTripImageUseCase(db *gorm.DB) TripImageUseCase {
	return &tripImageUseCase{
		db:           db,
		tripPlanRepo: repos.NewTripPlanRepo(db),
	}
}

func (tiu *tripImageUseCase) UploadFile(userID, tripID string, file *multipart.FileHeader) *common.AppError {
	tripPlan, err := tiu.tripPlanRepo.GetPlanByID(tripID)
	if err != nil {
		return common.InternalServerError(err)
	}
	if tripPlan == nil {
		return common.NotFoundError(errors.New("trip plan is nil"))
	}
	if tripPlan.UserID != userID {
		return common.BadRequestError(errors.New("user id does not match the expected value"))
	}

	src, err := file.Open()
	if err != nil {
		return common.InternalServerError(err)
	}
	defer src.Close()

	dstPath := fmt.Sprintf("tmp/upload/%s.png", tripID)
	dst, err := os.Create(dstPath)
	if err != nil {
		return common.InternalServerError(err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return common.InternalServerError(err)
	}

	return nil
}
