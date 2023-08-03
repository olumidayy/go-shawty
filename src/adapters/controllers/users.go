package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/domain/models"
	"github.com/olumidayy/go-shawty/src/usecases"
	"gorm.io/gorm"
)

type UserController struct {
	Usecases usecases.UserUsecases
}

func NewUserController (DB *gorm.DB) *UserController {
	return &UserController{
		Usecases: *usecases.NewUserUsecases(DB),
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User
	users = uc.Usecases.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	user := uc.Usecases.GetUserById(uint(id))
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not "})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	var user models.User
	user = uc.Usecases.UpdateUser(uint(id), c.Request.Body)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) DeleteUser (c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	uc.Usecases.DeleteUser(uint(id))
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}