package routes

import (
	"net/http"

	"tts-go-api/src/api"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api.SetupRoutes(r.Group("/api"))

	return r
}
