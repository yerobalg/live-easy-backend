package repository

import (
	firebase_auth "firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/src/entity"
)

type UserInterface interface {
	Get(ctx *gin.Context, params entity.UserParam) (entity.User, error)
	Create(ctx *gin.Context, user entity.User) (entity.User, error)
	VerifyFirebaseToken(ctx *gin.Context, firebaseJWT string) (*firebase_auth.Token, error)
}

type user struct {
	db       sql.DB
	firebase infrastructure.Firebase
}

func InitUser(db sql.DB, firebase infrastructure.Firebase) UserInterface {
	return &user{
		db:       db,
		firebase: firebase,
	}
}

func (u *user) VerifyFirebaseToken(ctx *gin.Context, firebaseJWT string) (*firebase_auth.Token, error) {
	return u.firebase.Auth.VerifyIDToken(ctx, firebaseJWT)
}

func (u *user) Get(ctx *gin.Context, params entity.UserParam) (entity.User, error) {
	var user entity.User

	res := u.db.GetDB(ctx).Where(params).First(&user)
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
