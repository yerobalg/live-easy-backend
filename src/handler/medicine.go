package handler

import (
	"github.com/gin-gonic/gin"
	"live-easy-backend/src/entity"
)

func (r *rest) CreateMedicine(ctx *gin.Context) {
	var medicineInput entity.MedicineInputParam
	if err := r.BindBody(ctx, &medicineInput); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	medicine, err := r.uc.Medicine.Create(ctx, medicineInput)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Successfully created medicine", medicine)
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

	SuccessResponse(ctx, "Successfully get medicine", medicine)
}
