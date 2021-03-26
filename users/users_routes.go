package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Signup struct {
	Email string
	Password string
}

func HandleUserRoutes(db *gorm.DB, r *gin.Engine) {

	r.GET("/users", func (c *gin.Context) {

		var users []User

		result := db.Find(&users)
		fmt.Println(result)

		c.JSON(200, gin.H{
			"message": users,
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		var user User
		db.First(&user, id)

		if user.ID != 0 {
			c.JSON(200, gin.H {
				"id": user,
			})
		} else {
			c.JSON(404, gin.H {
				"message": "No such user",
			})
		}

	})

	r.POST("/users", func (c *gin.Context) {

		var newUser Signup
		_ = c.Bind(&newUser)

		db.Create(&User{Email: newUser.Email, Password: newUser.Password})

		c.JSON(200, gin.H {
			"message": "User created for " + newUser.Email,
		})
	})

	r.DELETE("/users/:id", func (c *gin.Context) {
		var user User

		id := c.Param("id")

		db.First(&user, id)
		db.Delete(&user)

		c.JSON(200, gin.H {
			"message": "User Deleted",
		})
	})

	r.PUT("/users/:id", func (c *gin.Context) {
		var user User

		id := c.Param("id")

		var newUser Signup
		_ = c.Bind(&newUser)

		db.First(&user, id)
		db.Model(&user).Updates(User{Email: newUser.Email, Password: newUser.Password})

		c.JSON(200, gin.H {
			"message": "Updated",
		})
	})

	log.Fatal(r.Run(":8000"))
}
