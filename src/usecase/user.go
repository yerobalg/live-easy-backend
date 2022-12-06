package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/jwt"
	"live-easy-backend/sdk/password"
	"live-easy-backend/src/entity"
	"live-easy-backend/src/repository"
)

type UserInterface interface {
	Login(ctx *gin.Context, userParam entity.UserParam, userInput entity.UserLoginInputParam) (entity.UserLoginResponse, error)
	LoginWithGoogle(ctx *gin.Context, userGoogleInput entity.UserLoginGoogleInputParam) (entity.UserLoginResponse, error)
	Register(ctx *gin.Context, userInput entity.UserRegisterInputParam) (entity.User, error)
}

type User struct {
	userRepo repository.UserInterface
}

func InitUser(ur repository.UserInterface) UserInterface {
	return &User{
		userRepo: ur,
	}
}

func (u *User) Login(ctx *gin.Context, userParam entity.UserParam, userInput entity.UserLoginInputParam) (entity.UserLoginResponse, error) {
	var userResponse entity.UserLoginResponse

	userParam.Email = userInput.Email
	user, err := u.userRepo.Get(ctx, userParam)
	if err != nil {
		return userResponse, err
	}

	if user.IsGoogleAccount {
		return userResponse, errors.NewWithCode(401, "This is a google account, please login with google", "HTTPStatusUnauthorized")
	}

	if !password.Compare(user.Password, userInput.Password) {
		return userResponse, errors.NewWithCode(401, "Wrong password", "HTTPStatusUnauthorized")
	}

	token, err := jwt.GetToken(user)
	if err != nil {
		return userResponse, errors.NewWithCode(500, "Failed to generate token", "HTTPStatusInternalServerError")
	}

	userResponse.User = user
	userResponse.Token = token

	return userResponse, nil
}

func (u *User) LoginWithGoogle(ctx *gin.Context, userGoogleInput entity.UserLoginGoogleInputParam) (entity.UserLoginResponse, error) {
	var userResponse entity.UserLoginResponse

	firebaseUser, err := u.userRepo.GetFirebaseUser(ctx, userGoogleInput.FirebaseUID)
	if err != nil {
		return userResponse, err
	}

	userParam := entity.UserParam{Email: firebaseUser.Email}

	user, err := u.userRepo.Get(ctx, userParam)
	if errors.GetCode(err) == http.StatusNotFound {
		user, err = u.registerFromGoogleAccount(ctx, entity.UserRegisterInputParam{
			Email: firebaseUser.Email,
			Name:  firebaseUser.DisplayName,
		})
		if err != nil {
			return userResponse, err
		}
	} else if err != nil {
		return userResponse, err
	}

	token, err := jwt.GetToken(user)
	if err != nil {
		return userResponse, errors.NewWithCode(500, "Failed to generate token", "HTTPStatusInternalServerError")
	}

	userResponse.User = user
	userResponse.Token = token

	return userResponse, nil
}

func (u *User) registerFromGoogleAccount(ctx *gin.Context, userInput entity.UserRegisterInputParam) (entity.User, error) {
	var user entity.User

	user = entity.User{
		Email:           userInput.Email,
		Name:            userInput.Name,
		IsGoogleAccount: true,
	}

	user, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) Register(ctx *gin.Context, userInput entity.UserRegisterInputParam) (entity.User, error) {
	var user entity.User

	if !password.IsValid(userInput.Password) {
		return user, errors.NewWithCode(400, "Password must be at least 8 characters and contain at least 1 letter and 1 number", "HTTPStatusBadRequest")
	}

	hashedPassword, err := password.Hash(userInput.Password)
	if err != nil {
		return user, errors.NewWithCode(500, "Failed to hash password", "HTTPStatusInternalServerError")
	}

	user = entity.User{
		Email:    userInput.Email,
		Password: hashedPassword,
		Name:     userInput.Name,
	}

	user, err = u.userRepo.Create(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}
