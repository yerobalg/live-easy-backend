package repository

import (
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
)

type Repository struct {
	User UserInterface
}

func Init(db sql.DB, oauth infrastructure.OAuth) *Repository {
	return &Repository{
		User: InitUser(db, oauth),
	}
}
