package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olumidayy/go-shawty/src/infrastructure/database"
	"github.com/olumidayy/go-shawty/src/infrastructure/routes"
)

func main() {
  r := gin.Default()
	database.Connect()

	routes.SetupRoutes(r)

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run()
}
