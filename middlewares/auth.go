package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/OG-Chat-Backend/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Request"})
		return
	}
	email, err := utils.VerifyJWTToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	ctx.Set("email", email)
	ctx.Next()
}
