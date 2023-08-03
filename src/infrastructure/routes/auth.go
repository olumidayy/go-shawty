package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/adapters/controllers"
	"github.com/olumidayy/go-shawty/src/infrastructure/database"
)

func addAuthRoutes(rg *gin.Engine) {
	authRouter := rg.Group("/auth")
	authController := controllers.NewAuthController(database.DB)

	authRouter.POST("/register", authController.Register)
}