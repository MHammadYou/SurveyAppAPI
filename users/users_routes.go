package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type userInterface struct {
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

		var newUser userInterface
		_ = c.Bind(&newUser)

		hashPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil { panic(err) }

		//auth := bcrypt.CompareHashAndPassword(hashPass, []byte("fakePassword"))
		//
		//fmt.Println(auth)

		db.Create(&User{Email: newUser.Email, Password: string(hashPass)})

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

		var newUser userInterface
		_ = c.Bind(&newUser)

		db.First(&user, id)
		db.Model(&user).Updates(User{Email: newUser.Email, Password: newUser.Password})

		c.JSON(200, gin.H {
			"message": "Updated",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var loginData userInterface
		var user User

		err := c.Bind(&loginData)
		if err != nil { panic(err) }

		db.First(&user, "Email = ?", loginData.Email)

		if user.ID != 0 {
			login(loginData, user)
		} else {
			c.JSON(404, gin.H {
				"message": "No user exists with this email",
			})
		}

	})

	log.Fatal(r.Run(":8000"))
}


func login(loginData userInterface, user User) {

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err == nil {
		fmt.Println("Login Successful")
	} else {
		fmt.Println(err)
	}

}