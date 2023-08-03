package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(rg *gin.Engine) {
	addAuthRoutes(rg)
	addUserRoutes(rg)
}
