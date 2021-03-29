package main

import (
	"SurveyAppAPI/users"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {

	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&users.User{})

	handleRoutes(r, db)

	log.Fatal(r.Run(":8000"))

}
