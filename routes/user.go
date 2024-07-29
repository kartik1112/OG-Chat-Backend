package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/OG-Chat-Backend/models"
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
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Inavlid Credentials",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Logged In",
	})
}
