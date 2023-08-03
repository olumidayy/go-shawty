package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/domain/models"
	"github.com/olumidayy/go-shawty/src/usecases"
	"gorm.io/gorm"
)

type AuthController struct {
	Usecases usecases.UserUsecases
}

func NewAuthController (DB *gorm.DB) *AuthController {
	return &AuthController{
		Usecases: *usecases.NewUserUsecases(DB),
	}
}

type CreateUserInput struct {
	ID uint
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Email int `json:"email"`
  Password uint `json:"password"`
}

func (uc *AuthController) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	user := uc.Usecases.CreateNewUser(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}