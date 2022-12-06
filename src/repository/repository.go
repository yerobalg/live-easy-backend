package repository

import (
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
)

type Repository struct {
	User     UserInterface
	Medicine MedicineInterface
}

func Init(db sql.DB, firebase infrastructure.Firebase, storage infrastructure.Storage) *Repository {
	return &Repository{
		User:     InitUser(db, firebase),
		Medicine: InitMedicine(db, storage),
	}
}
