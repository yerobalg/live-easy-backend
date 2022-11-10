package repository

import (
	"github.com/gin-gonic/gin"
	"live-easy-backend/database/sql"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/src/entity"
)

type UserInterface interface {
	Get(ctx *gin.Context, params entity.UserParam) (entity.User, error)
	Create(ctx *gin.Context, user entity.User) (entity.User, error)
}

type user struct {
	db sql.DB
}

func InitUser(db sql.DB) UserInterface {
	return &user{db: db}
}

func (u *user) Get(ctx *gin.Context, params entity.UserParam) (entity.User, error) {
	var user entity.User

	whereClause := u.db.GetWhereClause(params)

	res := u.db.GetDB(ctx).Where(whereClause).First(&user)
	if res.RowsAffected == 0 {
		return user, errors.NotFound("User")
	} else if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

func (u *user) Create(ctx *gin.Context, user entity.User) (entity.User, error) {
	if err := u.db.GetDB(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
