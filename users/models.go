package users


import (
	"SurveyAppAPI/surveys"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Surveys []surveys.Survey `json:"surveys"`
}
