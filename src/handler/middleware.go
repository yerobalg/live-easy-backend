package handler

import (
	"github.com/gin-gonic/gin"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/jwt"
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
	ctx.Set("user", tokenClaims["data"])
	ctx.Next()
}

func (r *rest) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
