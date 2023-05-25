package v1

import (
	"net/http"

	"gurubear/models"
	"gurubear/repository/userrepo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type repository interface {
	CreateUser(user models.Users) bool
	GetUser() []models.Users
	UpdateUser(user models.Users) bool
	DeleteUser(ID uuid.UUID) bool
}

type UserController struct {
	//internal use
	service repository
}

// InitController initializes the user controller.
func UserInitController(userRepo *userrepo.UserRepo) *UserController {
	return &UserController{
		service: userRepo,
	}
}

func (controller *UserController) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, controller.service.GetUser())
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var user models.Users
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := controller.service.CreateUser(user)
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "success inerted"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail inserted"})
	}
}

func (controller *UserController) DeleteUser(c *gin.Context) {
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
	result := controller.service.DeleteUser(delID)
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "success deleted"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail deleted"})
	}

}
func (controller *UserController) UpdateUser(c *gin.Context) {
	var user models.Users
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := controller.service.UpdateUser(user)
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "success updated"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "fail updated"})
	}
}
