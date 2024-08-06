package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/OG-Chat-Backend/models"
)

func chatInit(ctx *gin.Context) {
	models.StartWs(&Hub, ctx.Writer, ctx.Request)
}

func checkServer(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"message": "wokring",
	})
}
