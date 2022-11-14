package repository

import (
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
)

type Repository struct {
	User     UserInterface
	Medicine MedicineInterface
}

func Init(db sql.DB, oauth infrastructure.OAuth, storage infrastructure.Storage) *Repository {
	return &Repository{
		User:     InitUser(db, oauth),
		Medicine: InitMedicine(db, storage),
	}
}
