package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/adapters/controllers"
	"github.com/olumidayy/go-shawty/src/adapters/repositories"
	"github.com/olumidayy/go-shawty/src/usecases"
	"gorm.io/gorm"
)

func addAuthRoutes(rg *gin.Engine, DB *gorm.DB) {
	authRouter := rg.Group("/auth")
	authController := controllers.NewAuthController(
		*usecases.NewAuthUsecases(
			repositories.NewUserRepository(DB),
		),
	)

	authRouter.POST("/register", authController.Register)
}