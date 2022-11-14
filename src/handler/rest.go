package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"live-easy-backend/infrastructure"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/src/usecase"
)

type rest struct {
	http  *gin.Engine
	uc    *usecase.Usecase
	oauth *infrastructure.OAuth
}

func Init(http *gin.Engine, uc *usecase.Usecase, o *infrastructure.OAuth) *rest {
	return &rest{
		http:  http,
		uc:    uc,
		oauth: o,
	}
}

func (r *rest) BindParam(ctx *gin.Context, param interface{}) error {
	if err := ctx.ShouldBindUri(param); err != nil {
		return err
	}

	return ctx.ShouldBindWith(param, binding.Query)
}

func (r *rest) BindBody(ctx *gin.Context, body interface{}) error {
	return ctx.ShouldBindWith(body, binding.Default(ctx.Request.Method, ctx.ContentType()))
}

type Response struct {
	Message   string      `json:"message"`
	IsSuccess bool        `json:"isSuccess"`
	Data      interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, Response{
		Message:   message,
		IsSuccess: true,
		Data:      data,
	})
}

func ErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(int(errors.GetCode(err)), Response{
		Message:   errors.GetType(err),
		IsSuccess: false,
		Data:      errors.GetMessage(err),
	})
}

func (r *rest) Run() {
	// Auth routes
	r.http.POST("api/v1/auth/register", r.Register)
	r.http.POST("api/v1/auth/login", r.Login)
	r.http.GET("api/v1/auth/login/google", r.LoginWithGoogle)
	r.http.GET("api/v1/auth/login/google/callback", r.LoginWithGoogleCallback)

	// Protected Routes
	v1 := r.http.Group("api/v1", r.Authorization())

	// User routes
	v1.Group("user")
	{
		v1.GET("user/profile", r.GetUserProfile)
	}

	// Medicine routes
	v1.Group("medicine")
	{
		v1.POST("medicine", r.CreateMedicine)
	}

	r.http.Run(":" + os.Getenv("APP_PORT"))
}
