package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func HandleUserRoutes(db *gorm.DB, r *gin.Engine) {

	r.GET("/users", handleGet)
	r.POST("/users", handlePost)

	log.Fatal(r.Run(":8000"))
}
