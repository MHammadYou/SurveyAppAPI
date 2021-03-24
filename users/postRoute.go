package users

import (
	"github.com/gin-gonic/gin"
)

type Signup struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func handlePost(c *gin.Context) {
	var signup Signup
	_ = c.BindJSON(&signup)
	c.JSON(200, gin.H {
		"message": signup,
	})
}