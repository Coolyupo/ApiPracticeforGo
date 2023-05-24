package main

import (
	"gurubear/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	r.POST("/createUser", func(ctx *gin.Context) {
		var user controller.Users
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := controller.CreateUser(user)
		if result {
			ctx.JSON(http.StatusOK, gin.H{"message": "success inerted"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "fail inserted"})
		}
	})

	r.POST("/deleteUser", func(ctx *gin.Context) {
		ID := ctx.PostForm("id")
		if ID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No Specific ID"})
			return
		}
		delID, err := uuid.Parse(ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := controller.DeleteUser(delID)
		if result {
			ctx.JSON(http.StatusOK, gin.H{"message": "success deleted"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "fail deleted"})
		}
	})
	r.POST("/updateUser", func(ctx *gin.Context) {
		var user controller.Users
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := controller.UpdateUser(user)
		if result {
			ctx.JSON(http.StatusOK, gin.H{"message": "success updated"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "fail updated"})
		}
	})

	r.Run(":8080")

}
