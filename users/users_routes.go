package users

import (
	"github.com/gin-gonic/gin"
	"log"
)

var r = gin.Default()

func Users() {
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Yo",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "Post Request",
		})
	})


	log.Fatal(r.Run(":8000"))
}