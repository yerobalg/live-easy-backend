package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserID(ctx *gin.Context) int64 {
	user := ctx.MustGet("user").(map[string]interface{})
	fmt.Println(user["id"].(float64))
	return int64(user["id"].(float64))
}
