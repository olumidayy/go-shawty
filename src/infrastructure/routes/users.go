package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/adapters/controllers"
	"github.com/olumidayy/go-shawty/src/infrastructure/database"
)

func addUserRoutes(rg *gin.Engine) {
	usersRouter := rg.Group("/users")
	userController := controllers.NewUserController(database.DB)

	usersRouter.GET("/", userController.GetAllUsers)
	usersRouter.GET("/:id", userController.GetUserByID)
	usersRouter.PATCH("/:id", userController.UpdateUser)
	usersRouter.DELETE("/", userController.DeleteUser)
}