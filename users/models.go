package users


import (
	"SurveyAppAPI/surveys"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique; not null;"`
	Password string `json:"password" gorm:"not null;"`
	Surveys []surveys.Survey `json:"surveys" gorm:"default=null;"`
}
