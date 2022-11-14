package usecase

import (
	"github.com/gin-gonic/gin"

	"live-easy-backend/sdk/auth"
	"live-easy-backend/sdk/null"
	"live-easy-backend/sdk/numeric"
	"live-easy-backend/src/entity"
	"live-easy-backend/src/repository"
)

type MedicineInterface interface {
	Create(ctx *gin.Context, medicineInput entity.MedicineInputParam) (entity.Medicine, error)
	Get(ctx *gin.Context, params entity.MedicineParam) (entity.Medicine, error)
	GetList(ctx *gin.Context, params entity.MedicineParam) ([]entity.Medicine, *entity.PaginationParam, error)
	Update(ctx *gin.Context, medicineParam entity.MedicineParam) error
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

func (m *Medicine) Create(ctx *gin.Context, medicineInput entity.MedicineInputParam) (entity.Medicine, error) {
	userID := auth.GetUserID(ctx)
	medicine := entity.Medicine{
		Name:        medicineInput.Name,
		Price:       medicineInput.Price,
		PriceString: numeric.IntToRupiah(medicineInput.Price),
		Quantity:    medicineInput.Quantity,
		UserID:      userID,
		CreatedBy:   null.Int64From(userID),
		UpdatedBy:   null.Int64From(userID),
	}

	medicine, err := m.repo.Create(ctx, medicine)
	if err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (m *Medicine) Get(ctx *gin.Context, params entity.MedicineParam) (entity.Medicine, error) {
	medicine, err := m.repo.Get(ctx, params)
	if err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (m *Medicine) GetList(ctx *gin.Context, params entity.MedicineParam) ([]entity.Medicine, *entity.PaginationParam, error) {
	medicines, pg, err := m.repo.GetList(ctx, params)
	if err != nil {
		return medicines, pg, err
	}

	return medicines, pg, nil
}

func (m *Medicine) Update(ctx *gin.Context, medicineParam entity.MedicineParam) error {
	err := m.repo.Update(ctx, medicineParam)
	if err != nil {
		return err
	}

	return nil
}

func (m *Medicine) Delete(ctx *gin.Context, medicineParam entity.MedicineParam) error {
	err := m.repo.Delete(ctx, medicineParam)
	if err != nil {
		return err
	}

	return nil
}
