package handler

import (
	"fmt"
	"time"

	"live-easy-backend/sdk/appcontext"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/src/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (r *rest) BindParam(ctx *gin.Context, param interface{}) error {
	if err := ctx.ShouldBindUri(param); err != nil {
		return err
	}

	return ctx.ShouldBindWith(param, binding.Query)
}

func (r *rest) BindBody(ctx *gin.Context, body interface{}) error {
	return ctx.ShouldBindWith(body, binding.Default(ctx.Request.Method, ctx.ContentType()))
}

func getRequestMetadata(ctx *gin.Context) entity.Meta {
	meta := entity.Meta{
		RequestID: appcontext.GetRequestId(ctx),
		Time:      time.Now().Format(time.RFC3339),
	}

	requestStartTime := appcontext.GetRequestStartTime(ctx)
	if !requestStartTime.IsZero() {
		elapsedTimeMs := time.Since(requestStartTime).Milliseconds()
		meta.TimeElapsed = fmt.Sprintf("%dms", elapsedTimeMs)
	}

	return meta
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}, pg *entity.PaginationParam) {
	ctx.JSON(200, entity.HTTPResponse{
		Meta :      getRequestMetadata(ctx),
		Message:    message,
		IsSuccess:  true,
		Data:       data,
		Pagination: pg,
	})
}

func ErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(int(errors.GetCode(err)), entity.HTTPResponse{
		Meta:      getRequestMetadata(ctx),
		Message:   errors.GetType(err),
		IsSuccess: false,
		Data:      errors.GetMessage(err),
	})
}
