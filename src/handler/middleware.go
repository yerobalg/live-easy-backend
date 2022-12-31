package handler

import (
	"context"
	"time"

	"live-easy-backend/sdk/appcontext"
	"live-easy-backend/sdk/errors"
	"live-easy-backend/sdk/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// timeout middleware wraps the request context with a timeout
func (r *rest) SetTimeout(ctx *gin.Context) {
	// wrap the request context with a timeout, this will cause the request to fail if it takes more than defined timeout
	c, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Minute) // TODO: change this hardcoded timeout to config later

	// cancel to clear resources after finished
	defer cancel()

	appcontext.SetRequestStartTime(ctx, time.Now())

	// replace request with context wrapped request
	ctx.Request = ctx.Request.WithContext(c)
	ctx.Next()
}

func (r *rest) AddFieldsToContext(ctx *gin.Context) {
	requestID := uuid.New().String()

	appcontext.SetRequestId(ctx, requestID)
	appcontext.SetUserAgent(ctx, ctx.Request.Header.Get(appcontext.HeaderUserAgent))
	appcontext.SetDeviceType(ctx, ctx.Request.Header.Get(appcontext.HeaderDeviceType))

	ctx.Next()
}

func (r *rest) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
