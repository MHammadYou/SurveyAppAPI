package surveys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleSurveysRoutes(db *gorm.DB, r *gin.Engine) {

	r.GET("/surveys/", func(c *gin.Context) {
		var surveys []Survey

		db.Find(&surveys)

		c.JSON(200, gin.H{
			"surveys": surveys,
		})
	})

	r.GET("/surveys/:id", func(c *gin.Context) {

		id := c.Param("id")

		var survey Survey
		db.First(&survey, id)

		if survey.ID != 0 {
			c.JSON(200, gin.H {
				"survey": survey,
			})
		} else {
			c.JSON(404, gin.H {
				"message": "No such survey",
			})
		}
	})

}
