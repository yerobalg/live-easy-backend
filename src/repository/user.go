package repository

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/src/entity"
)

type UserInterface interface {
	Get(ctx *gin.Context, params entity.UserParam) (entity.User, error)
	Create(ctx *gin.Context, user entity.User) (entity.User, error)
	GoogleCallback(ctx *gin.Context, code string) (map[string]interface{}, error)
}

type user struct {
	db    sql.DB
	oauth infrastructure.OAuth
}

func InitUser(db sql.DB, oauth infrastructure.OAuth) UserInterface {
	return &user{
		db:    db,
		oauth: oauth,
	}
}

func (u *user) GoogleCallback(ctx *gin.Context, code string) (map[string]interface{}, error) {
	var userResponseGoogle map[string]interface{}

	googleConfig := u.oauth.Config
	token, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		return userResponseGoogle, errors.NewWithCode(500, err.Error(), "HTTPStatusInternalServerError")
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return userResponseGoogle, errors.NewWithCode(500, err.Error(), "HTTPStatusInternalServerError")
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return userResponseGoogle, errors.NewWithCode(500, err.Error(), "HTTPStatusInternalServerError")
	}

	err = json.Unmarshal(responseData, &userResponseGoogle)
	if err != nil {
		return userResponseGoogle, errors.NewWithCode(500, err.Error(), "HTTPStatusInternalServerError")
	}

	return userResponseGoogle, nil
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
