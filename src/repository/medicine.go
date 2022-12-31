package repository

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/file"
	"live-easy-backend/src/entity"
)

type MedicineInterface interface {
	Create(ctx context.Context, medicine entity.Medicine) (entity.Medicine, error)
	Get(ctx context.Context, params entity.MedicineParam) (entity.Medicine, error)
	GetList(ctx context.Context, params entity.MedicineParam) ([]entity.Medicine, *entity.PaginationParam, error)
	Update(ctx context.Context, medicineParam entity.MedicineParam, medicine entity.Medicine) error
	Delete(ctx context.Context, medicineParam entity.MedicineParam) error
	UploadImage(ctx context.Context, file *file.File) (string, error)
	DeleteImage(ctx context.Context, imageURL string) error
}

type medicine struct {
	db      sql.DB
	storage infrastructure.Storage
}

func InitMedicine(db sql.DB, storage infrastructure.Storage) MedicineInterface {
	return &medicine{
		db:      db,
		storage: storage,
	}
}

func (m *medicine) Create(ctx context.Context, medicine entity.Medicine) (entity.Medicine, error) {
	if err := m.db.WithContext(ctx).Create(&medicine).Error; err != nil {
		return medicine, err
	}

	return medicine, nil
}

func (m *medicine) Get(ctx context.Context, params entity.MedicineParam) (entity.Medicine, error) {
	var medicine entity.Medicine

	res := m.db.WithContext(ctx).Where(params).First(&medicine)
	if res.RowsAffected == 0 {
		return medicine, errors.NotFound("Medicine")
	} else if res.Error != nil {
		return medicine, res.Error
	}

	return medicine, nil
}

func (m *medicine) GetList(ctx context.Context, params entity.MedicineParam) ([]entity.Medicine, *entity.PaginationParam, error) {
	var medicine []entity.Medicine

	pg := entity.FormatPaginationParam(params.PaginationParam)

	if err := m.db.
		WithContext(ctx).
		Where(params).
		Offset(int(pg.Offset)).
		Limit(int(pg.Limit)).
		Find(&medicine).Error; err != nil {
		return medicine, nil, err
	}

	if err := m.db.
		WithContext(ctx).
		Model(&medicine).
		Where(params).
		Count(&pg.TotalElement).Error; err != nil {
		return medicine, nil, err
	}
	pg.ProcessPagination(int64(len(medicine)))

	return medicine, &pg, nil
}

func (m *medicine) Update(ctx context.Context, medicineParam entity.MedicineParam, medicine entity.Medicine) error {
	res := m.db.
		WithContext(ctx).
		Model(entity.Medicine{}).
		Where(&medicineParam).
		Updates(&medicine)

	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected == 0 {
		return errors.NotFound("Medicine")
	}

	return nil
}

func (m *medicine) Delete(ctx context.Context, medicineParam entity.MedicineParam) error {
	res := m.db.
		WithContext(ctx).
		Where(&medicineParam).
		Delete(&entity.Medicine{})
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected == 0 {
		return errors.NotFound("Medicine")
	}

	return nil
}

func (m *medicine) UploadImage(ctx context.Context, file *file.File) (string, error) {
	var imageURL string

	storageHandler, err := m.storage.GetObjectPlace(fmt.Sprintf("%s/%s", m.storage.FolderName, file.Meta.Filename))
	if err != nil {
		return imageURL, errors.InternalServerError(err.Error())
	}

	storageWriter := storageHandler.NewWriter(ctx)

	if _, err := io.Copy(storageWriter, file.Content); err != nil {
		return imageURL, errors.InternalServerError(err.Error())
	}

	if err := storageWriter.Close(); err != nil {
		return imageURL, errors.InternalServerError(err.Error())
	}

	parsedURL, err := url.Parse(fmt.Sprintf(
		"%s/%s/%s/%s",
		m.storage.BaseURL,
		m.storage.BucketName,
		m.storage.FolderName,
		file.Meta.Filename,
	))
	if err != nil {
		return imageURL, errors.InternalServerError(err.Error())
	}
	imageURL = parsedURL.String()

	return imageURL, nil
}

func (m *medicine) DeleteImage(ctx context.Context, fileName string) error {
	storageHandler, err := m.storage.GetObjectPlace(fmt.Sprintf("%s/%s", m.storage.FolderName, fileName))
	if err != nil {
		return err
	}

	if err = storageHandler.Delete(ctx); err != nil {
		return err
	}

	return nil
}
