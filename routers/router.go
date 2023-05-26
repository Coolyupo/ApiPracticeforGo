package routers

import (
	"gurubear/controller"
	"gurubear/database"
	"gurubear/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controller := controller.InitControllers(repos)

	apiv1 := r.Group("/api/v1")

	apiv1.POST("/createUser", controller.UserController.CreateUser)
	apiv1.GET("/getUser", controller.UserController.GetUser)
	apiv1.POST("/deleteUser", controller.UserController.DeleteUser)
	apiv1.POST("/updateUser", controller.UserController.UpdateUser)

	apiv1.POST("/createPost", controller.PostController.CreatePost)
	apiv1.GET("/getPost", controller.PostController.GetPost)
	apiv1.POST("/updatePost", controller.PostController.UpdatePost)
	apiv1.POST("/deletePost", controller.PostController.DeletePost)

	return r
}
