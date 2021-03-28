package surveys

import "gorm.io/gorm"

type Survey struct {
	gorm.Model
	Title string
	Question string
	Answer string
}

