package sql

import (
	"live-easy-backend/src/entity"
	"gorm.io/gorm"
)

type Migration struct {
	Db *gorm.DB
}

func (m *Migration) RunMigration() {
	m.Db.AutoMigrate(
		&entity.User{},
	)
}

