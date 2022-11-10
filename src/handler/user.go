package handler

import (
	"github.com/gin-gonic/gin"
	"live-easy-backend/src/entity"
)

func (r *rest) Login(ctx *gin.Context) {
	var userParam entity.UserParam
	var userInput entity.UserLoginInputParam

	if err := r.BindParam(ctx, &userParam); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	if err := r.BindBody(ctx, &userInput); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	userResponse, err := r.uc.User.Login(ctx, userParam, userInput)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Login success", userResponse)
}

func (r *rest) GetUserProfile(ctx *gin.Context) {
	user := ctx.MustGet("user")

	SuccessResponse(ctx, "Get user profile success", user)
}

func (r *rest) Register(ctx *gin.Context) {
	var userInput entity.UserRegisterInputParam

	if err := r.BindBody(ctx, &userInput); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	user, err := r.uc.User.Register(ctx, userInput)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Register success", user)
}
