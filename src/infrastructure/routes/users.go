package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/adapters/controllers"
	"github.com/olumidayy/go-shawty/src/adapters/repositories"
	"github.com/olumidayy/go-shawty/src/usecases"
	"gorm.io/gorm"
)

func addUserRoutes(rg *gin.Engine, DB *gorm.DB) {
	usersRouter := rg.Group("/users")
	userController := controllers.NewUserController(
		*usecases.NewUserUsecases(
			repositories.NewUserRepository(DB),
		))

	usersRouter.GET("/", userController.GetAllUsers)
	usersRouter.GET("/:id", userController.GetUserByID)
	usersRouter.PATCH("/:id", userController.UpdateUser)
	usersRouter.DELETE("/:id", userController.DeleteUser)
}
