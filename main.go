package main

import (
	"gurubear/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	// o := controller.GetUser()
	r.GET("/getUser", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.GetUser())
	})

	r.Run(":8080")

}
