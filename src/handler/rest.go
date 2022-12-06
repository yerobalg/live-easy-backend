package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/src/entity"
	"live-easy-backend/src/usecase"
)

type rest struct {
	http *gin.Engine
	uc   *usecase.Usecase
}

func Init(http *gin.Engine, uc *usecase.Usecase) *rest {
	return &rest{
		http: http,
		uc:   uc,
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
	Message    string                  `json:"message"`
	IsSuccess  bool                    `json:"isSuccess"`
	Data       interface{}             `json:"data"`
	Pagination *entity.PaginationParam `json:"pagination"`
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}, pg *entity.PaginationParam) {
	ctx.JSON(200, Response{
		Message:    message,
		IsSuccess:  true,
		Data:       data,
		Pagination: pg,
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
	r.http.Use(r.CorsMiddleware())
	// Auth routes
	r.http.POST("api/v1/auth/register", r.Register)
	r.http.POST("api/v1/auth/login", r.Login)
	r.http.POST("api/v1/auth/login/google", r.LoginWithGoogle)

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
		v1.GET("medicine/:id", r.GetMedicine)
		v1.GET("medicine", r.GetListMedicines)
		v1.PUT("medicine/:id", r.UpdateMedicine)
		v1.DELETE("medicine/:id", r.DeleteMedicine)
	}

	r.http.Run(":" + os.Getenv("APP_PORT"))
}
