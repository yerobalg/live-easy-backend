package repository

import (
	"live-easy-backend/database/sql"
)

type Repository struct {
	User UserInterface
}

func Init(db sql.DB) *Repository {
	return &Repository{
		User: InitUser(db),
	}
}
