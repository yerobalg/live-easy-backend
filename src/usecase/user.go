package usecase

import (
	"context"
	"net/http"

	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/jwt"
	"live-easy-backend/sdk/password"
	"live-easy-backend/src/entity"
	"live-easy-backend/src/repository"
)

type UserInterface interface {
	Login(ctx context.Context, userParam entity.UserParam, userInput entity.UserLoginInputParam) (entity.UserLoginResponse, error)
	LoginWithGoogle(ctx context.Context, userGoogleInput entity.UserLoginGoogleInputParam) (entity.UserLoginResponse, error)
	Register(ctx context.Context, userInput entity.UserRegisterInputParam) (entity.User, error)
}

type User struct {
	userRepo repository.UserInterface
}

func InitUser(ur repository.UserInterface) UserInterface {
	return &User{
		userRepo: ur,
	}
}

func (u *User) Login(ctx context.Context, userParam entity.UserParam, userInput entity.UserLoginInputParam) (entity.UserLoginResponse, error) {
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

func (u *User) LoginWithGoogle(ctx context.Context, userGoogleInput entity.UserLoginGoogleInputParam) (entity.UserLoginResponse, error) {
	var userResponse entity.UserLoginResponse

	firebaseToken, err := u.userRepo.VerifyFirebaseToken(ctx, userGoogleInput.FirebaseJWT)
	if err != nil {
		return userResponse, errors.NewWithCode(401, "Invalid firebase token", "HTTPStatusUnauthorized")
	}

	email := firebaseToken.Claims["email"].(string)
	name := firebaseToken.Claims["name"].(string)

	user, err := u.userRepo.Get(ctx, entity.UserParam{Email: email, IsGoogleAccount: true})
	if errors.GetCode(err) == http.StatusNotFound {
		user, err = u.registerFromGoogleAccount(ctx, entity.UserRegisterInputParam{
			Email: email,
			Name:  name,
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

func (u *User) registerFromGoogleAccount(ctx context.Context, userInput entity.UserRegisterInputParam) (entity.User, error) {
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

func (u *User) Register(ctx context.Context, userInput entity.UserRegisterInputParam) (entity.User, error) {
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
