package handlers

import (
	t "auth-service/api/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		url := ctx.Request.URL.Path

		if strings.Contains(url, "swagger") || url == "/v1/login" {
			ctx.Next()
			return
		}

		isValid, err := t.Validate(token)
		if err != nil || !isValid {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid or expired token"})
			return
		}

		ctx.Next()
	}
}
