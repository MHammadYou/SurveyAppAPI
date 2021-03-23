package main

import (
	"SurveyAppAPI/users"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	users.Users()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&users.User{})

	var user1 users.User

	db.First(&user1, 1)

	db.Model(&user1).Update("Email", "TestEmail@testing.com")

	fmt.Printf("%+v\n", user1)
}
