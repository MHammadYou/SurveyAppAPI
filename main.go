package main

import (
	"SurveyAppAPI/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {


	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&users.User{})

	handleRoutes(db)

}
