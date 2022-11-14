package repository

import (
	"github.com/gin-gonic/gin"
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
	"live-easy-backend/src/entity"
)

type MedicineIntercafe interface {
}

type medicine struct {
	db      sql.DB
	storage infrastructure.Storage
}

func InitMedicine(db sql.DB, storage infrastructure.Storage) MedicineIntercafe {
	return &medicine{
		db:      db,
		storage: storage,
	}
}

func (m *medicine) Create(ctx *gin.Context, medicine entity.Medicine) (entity.Medicine, error) {
	if err := m.db.Create(&medicine).Error; err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (m *medicine) Get(ctx *gin.Context, params entity.MedicineParam) (entity.Medicine, error) {
	var medicine entity.Medicine

	if err := m.db.Where(params).First(&medicine).Error; err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (m *medicine) GetList(ctx *gin.Context, params entity.MedicineParam) ([]entity.Medicine, *entity.PaginationParam, error) {
	var medicine []entity.Medicine

	pg := params.PaginationParam
	offset := pg.GetOffset()

	if err := m.db.
		Where(params).
		Offset(int(offset)).
		Limit(int(pg.Limit)).
		Find(&medicine).Error; err != nil {
		return medicine, nil, err
	}

	if err := m.db.
		Model(&medicine).
		Where(params).
		Count(&pg.TotalElement).Error; err != nil {
		return medicine, nil, err
	}
	pg.ProcessPagination()

	return medicine, &pg, nil
}
