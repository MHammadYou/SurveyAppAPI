package surveys

import "gorm.io/gorm"

type Survey struct {
	gorm.Model
	Title string `json:"title"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}

