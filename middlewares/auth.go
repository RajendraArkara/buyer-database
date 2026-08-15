package middlewares

import (
	"net/http"
	"strings"

	"github.com/RajendraArkara/buyer-database/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

func Authencticate(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}
