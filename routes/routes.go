package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/OG-Chat-Backend/db"
	"github.com/kartik1112/OG-Chat-Backend/middlewares"
	"github.com/kartik1112/OG-Chat-Backend/models"
)

var Hub models.Hub

func RegisterRoutes(server *gin.Engine) {
	Hub.NewHub()
	go Hub.Run()
	server.POST("/api/users/register", signup)
	server.POST("/api/users/login", login)
	Authenticated := server.Group("/")
	Authenticated.Use(middlewares.Authenticate)
	Authenticated.GET("/api/status", dbHealthCheck)
	Authenticated.GET("/api/users/", getUserByEmail)
	Authenticated.PUT("/api/users/", updateUserByEmail)
	// server.GET("/echo", checkServer)
	Authenticated.GET("/ws", chatInit)
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
