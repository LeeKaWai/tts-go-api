package api

import (
	"net/http"

	"tts-go-api/src/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/user/:id", getUser)
}

func getUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := service.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
