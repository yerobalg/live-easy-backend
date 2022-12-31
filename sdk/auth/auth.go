package auth

import (
	"context"
)

type key string

const (
	userAuthInfo key = "UserAuthInfo"
)

type User struct {
	ID              int64  `json:"id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	IsGoogleAccount bool   `json:"isGoogleAccount"`
}

func GetUserID(ctx context.Context) int64 {
	user, ok := ctx.Value(userAuthInfo).(User)
	if !ok {
		return 0
	}

	return user.ID
}

func GetUser(ctx context.Context) User {
	user, ok := ctx.Value(userAuthInfo).(User)
	if !ok {
		return User{}
	}

	return user
}

func SetUser(ctx context.Context, user map[string]interface{}) context.Context {
	userObj := User{
		ID:              int64(user["id"].(float64)),
		Email:           user["email"].(string),
		Name:            user["name"].(string),
		IsGoogleAccount: user["isGoogleAccount"].(bool),
	}
	return context.WithValue(ctx, userAuthInfo, userObj)
}
