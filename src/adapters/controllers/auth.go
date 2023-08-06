package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/domain/models"
	"github.com/olumidayy/go-shawty/src/usecases"
)

type AuthController struct {
	Usecases usecases.AuthUsecases
}

func NewAuthController (au usecases.AuthUsecases) *AuthController {
	return &AuthController{
		Usecases: au,
	}
}


func (uc *AuthController) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	user, err := uc.Usecases.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}