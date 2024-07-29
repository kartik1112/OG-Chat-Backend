package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/OG-Chat-Backend/db"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/status", dbHealthCheck)
	server.POST("/api/users/register", signup)
	server.POST("/api/users/login", login)
}

func dbHealthCheck(ctx *gin.Context) {
	err := db.DB.Ping()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not connect to DB",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DB connected All Systems Working Fine",
	})
}
