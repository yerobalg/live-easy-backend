package usecase

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"live-easy-backend/sdk/auth"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/file"
	"live-easy-backend/sdk/null"
	"live-easy-backend/sdk/numeric"
	"live-easy-backend/src/entity"
	"live-easy-backend/src/repository"
)

type MedicineInterface interface {
	Create(ctx *gin.Context, medicineInput entity.MedicineInputParam, image *file.File) (entity.Medicine, error)
	Get(ctx *gin.Context, params entity.MedicineParam) (entity.Medicine, error)
	GetList(ctx *gin.Context, params entity.MedicineParam) ([]entity.Medicine, *entity.PaginationParam, error)
	Update(ctx *gin.Context, medicineParam entity.MedicineParam, medicineInput entity.MedicineUpdateInputParam, image *file.File) error
	Delete(ctx *gin.Context, medicineParam entity.MedicineParam) error
}

type Medicine struct {
	repo repository.MedicineInterface
}

func InitMedicine(repo repository.MedicineInterface) MedicineInterface {
	return &Medicine{
		repo: repo,
	}
}

func (m *Medicine) Create(ctx *gin.Context, medicineInput entity.MedicineInputParam, image *file.File) (entity.Medicine, error) {
	userID := auth.GetUserID(ctx)
	now := time.Now().Unix()
	medicine := entity.Medicine{
		Name:        medicineInput.Name,
		Price:       medicineInput.Price,
		PriceString: numeric.IntToRupiah(medicineInput.Price),
		Quantity:    medicineInput.Quantity,
		UserID:      userID,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   null.Int64From(userID),
		UpdatedBy:   null.Int64From(userID),
	}

	if !image.IsImage() {
		return medicine, errors.BadRequest("file is not an image")
	}

	image.SetFileName(fmt.Sprintf("%d_%d", userID, now))

	imageURL, err := m.repo.UploadImage(ctx, image)
	if err != nil {
		return medicine, err
	}

	medicine.ImageURL = imageURL

	medicine, err = m.repo.Create(ctx, medicine)
	if err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (m *Medicine) Get(ctx *gin.Context, params entity.MedicineParam) (entity.Medicine, error) {
	params.UserID = auth.GetUserID(ctx)

	medicine, err := m.repo.Get(ctx, params)
	if err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (m *Medicine) GetList(ctx *gin.Context, params entity.MedicineParam) ([]entity.Medicine, *entity.PaginationParam, error) {
	params.UserID = auth.GetUserID(ctx)

	medicines, pg, err := m.repo.GetList(ctx, params)
	if err != nil {
		return medicines, pg, err
	}

	return medicines, pg, nil
}

func (m *Medicine) Update(ctx *gin.Context, medicineParam entity.MedicineParam, medicineInput entity.MedicineUpdateInputParam, image *file.File) error {
	userID := auth.GetUserID(ctx)

	medicineParam.UserID = userID

	var medicine entity.Medicine
	var imageURL string

	if image != nil {
		medicine, err := m.repo.Get(ctx, medicineParam)
		if err != nil {
			return err
		}

		image.SetFileName(fmt.Sprintf("%d_%d", medicine.UserID, medicine.CreatedAt))
		imageURL, err = m.repo.UploadImage(ctx, image)
		if err != nil {
			return err
		}
	}

	medicine = entity.Medicine{
		Name:      medicineInput.Name,
		Price:     medicineInput.Price,
		Quantity:  medicineInput.Quantity,
		UpdatedBy: null.Int64From(userID),
		ImageURL:  imageURL,
	}

	if medicine.Price != int64(0) {
		medicine.PriceString = numeric.IntToRupiah(medicine.Price)
	}

	err := m.repo.Update(ctx, medicineParam, medicine)
	if err != nil {
		return err
	}

	return nil
}

func (m *Medicine) Delete(ctx *gin.Context, medicineParam entity.MedicineParam) error {
	medicineParam.UserID = auth.GetUserID(ctx)

	err := m.repo.Delete(ctx, medicineParam)
	if err != nil {
		return err
	}

	return nil
}
