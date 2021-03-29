package main

import (
	"SurveyAppAPI/surveys"
	"SurveyAppAPI/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func handleRoutes(r *gin.Engine, db *gorm.DB) {

	users.HandleUserRoutes(db, r)
	surveys.HandleSurveysRoutes(db, r)

}
