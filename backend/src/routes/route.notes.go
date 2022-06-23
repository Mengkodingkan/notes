package route

import (
	model "mengkodingkan/notes/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitNotesRoutes(db *gorm.DB, route *gin.Engine) {
	/*
		Group Route
	*/

	group := route.Group("/api/v1")

	group.GET("/notes", func(c *gin.Context) {
		var users []model.Note

		datas := db.Find(&users)
		if datas.Error == nil {
			c.JSON(200, gin.H{
				"status": "success",
				"data":   users,
			})
		} else {
			// send data
			c.JSON(200, gin.H{
				"status": "success",
			})
		}
	})
}
