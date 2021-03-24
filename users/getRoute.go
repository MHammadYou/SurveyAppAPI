package users

import "github.com/gin-gonic/gin"

func handleGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Yo",
	})
}

