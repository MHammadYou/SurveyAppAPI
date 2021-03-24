package main

import (
	"SurveyAppAPI/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func handleRoutes(db *gorm.DB) {

	var r = gin.Default()

	users.HandleUserRoutes(db, r)
}
