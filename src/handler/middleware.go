package handler

import (
	"fmt"

	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/jwt"
	"github.com/gin-gonic/gin"
)

func (r *rest) Authorization() gin.HandlerFunc {
	return r.checkToken
}

func (r *rest) checkToken(ctx *gin.Context) {
	header := ctx.Request.Header.Get("Authorization")
	if header == "" {
		ErrorResponse(ctx, errors.NewWithCode(401, "Unauthorized", "Please login first"))
		ctx.Abort()
		return
	}

	header = header[len("Bearer "):]
	tokenClaims, err := jwt.DecodeToken(header)
	if err != nil {
		ErrorResponse(ctx, err)
		ctx.Abort()
		return
	}

	fmt.Println(tokenClaims["data"])

	ctx.Set("user", tokenClaims["data"])
	ctx.Next()
	return
}
