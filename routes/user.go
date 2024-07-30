package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/OG-Chat-Backend/models"
	"github.com/kartik1112/OG-Chat-Backend/utils"
)

func signup(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBindJSON(&user)
	fmt.Println(user)
	err := user.CreateUser()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Sign Up failed",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Signed Up",
	})
}

func login(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBindJSON(&user)
	err := user.ValidateUser()
	fmt.Print(user.UserID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Inavlid Credentials",
		})
		return
	}
	token, err := utils.GenerateJWTToken(user.Email, user.UserID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Inavlid Credentials",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Logged In",
		"token":   token,
	})

}
