package handler

import (
	"github.com/gin-gonic/gin"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/file"
	"live-easy-backend/src/entity"
)

func (r *rest) CreateMedicine(ctx *gin.Context) {
	var medicineInput entity.MedicineInputParam
	if err := r.BindBody(ctx, &medicineInput); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	image, err := file.Init(ctx, "image")
	if err != nil {
		ErrorResponse(ctx, errors.BadRequest(err.Error()))
		return
	}

	medicine, err := r.uc.Medicine.Create(ctx, medicineInput, image)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Successfully created medicine", medicine, nil)
}

func (r *rest) GetMedicine(ctx *gin.Context) {
	var medicineParam entity.MedicineParam
	if err := r.BindParam(ctx, &medicineParam); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	medicine, err := r.uc.Medicine.Get(ctx, medicineParam)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Successfully get medicine", medicine, nil)
}

func (r *rest) GetListMedicines(ctx *gin.Context) {
	var medicineParam entity.MedicineParam
	if err := r.BindParam(ctx, &medicineParam); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	medicine, pg, err := r.uc.Medicine.GetList(ctx, medicineParam)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Successfully get list of medicine", medicine, pg)
}

func (r *rest) UpdateMedicine(ctx *gin.Context) {
	var medicineParam entity.MedicineParam
	if err := r.BindParam(ctx, &medicineParam); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	var medicineInput entity.MedicineUpdateInputParam
	if err := r.BindBody(ctx, &medicineInput); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	image, _ := file.Init(ctx, "image")

	err := r.uc.Medicine.Update(ctx, medicineParam, medicineInput, image)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Successfully updated medicine", nil, nil)
}

func (r *rest) DeleteMedicine(ctx *gin.Context) {
	var medicineParam entity.MedicineParam
	if err := r.BindParam(ctx, &medicineParam); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err := r.uc.Medicine.Delete(ctx, medicineParam)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Successfully deleted medicine", nil, nil)
}
