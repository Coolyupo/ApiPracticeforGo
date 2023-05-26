package post

import (
	"gurubear/models"
	"gurubear/repository/postrepo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type repository interface {
	GetPost() []models.Posts
	CreatePost(post models.Posts) bool
	UpdatePost(post models.Posts) bool
	DeletePost(ID uuid.UUID) bool
}

type Controller struct {
	//internal use
	service repository
}

func InitController(postRepo *postrepo.PostRepo) *Controller {
	return &Controller{
		service: postRepo,
	}
}

func (controller *Controller) GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, controller.service.GetPost())
}

func (controller *Controller) CreatePost(c *gin.Context) {
	var post models.Posts
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := controller.service.CreatePost(post)
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "success inerted"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail inserted"})
	}
}

func (controller *Controller) UpdatePost(c *gin.Context) {
	var post models.Posts
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := controller.service.UpdatePost(post)
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "success updated"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "fail updated"})
	}
}

func (controller *Controller) DeletePost(c *gin.Context) {
	ID := c.PostForm("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Specific ID"})
		return
	}
	delID, err := uuid.Parse(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := controller.service.DeletePost(delID)
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "success deleted"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail deleted"})
	}

}
