package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/domain/models"
	"github.com/olumidayy/go-shawty/src/usecases"
)

type UserController struct {
	Usecases usecases.UserUsecases
}

func NewUserController (uu usecases.UserUsecases) *UserController {
	return &UserController{
		Usecases: uu,
	}
}


func (uc *UserController) GetAllUsers(c *gin.Context) {
	var users *[]models.User
	users, err := uc.Usecases.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}


func (uc *UserController) GetUserByID(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	user, err := uc.Usecases.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}


func (uc *UserController) UpdateUser(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)

	var data *models.User
	if err := c.ShouldBindJSON(&data); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	var user *models.User
	user, err = uc.Usecases.UpdateUser(uint(id), data)
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