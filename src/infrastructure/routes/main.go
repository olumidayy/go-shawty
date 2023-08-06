package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(rg *gin.Engine, DB *gorm.DB) {
	addAuthRoutes(rg, DB)
	addUserRoutes(rg, DB)
}
