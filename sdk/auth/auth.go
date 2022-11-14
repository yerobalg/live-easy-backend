package auth

import (
	"github.com/gin-gonic/gin"
)

func GetUserID(ctx *gin.Context) int64 {
	user := ctx.MustGet("user").(map[string]interface{})
	return int64(user["id"].(float64))
}
