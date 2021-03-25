package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Signup struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func HandleUserRoutes(db *gorm.DB, r *gin.Engine) {

	r.GET("/users", func (c *gin.Context) {

		var users []User

		db.Create(&User{Email: "jd@testing.com", Password: "testing1"})

		result := db.Find(&users)
		fmt.Println(result)

		c.JSON(200, gin.H{
			"message": users,
		})
	})

	r.POST("/users", func (c *gin.Context) {
		var signup Signup
		_ = c.Bind(&signup)
		c.JSON(200, gin.H {
			"message": signup,
		})
	})

	log.Fatal(r.Run(":8000"))
}
