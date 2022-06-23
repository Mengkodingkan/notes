package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitNotesRoutes(db *gorm.DB, route *gin.Engine) {
	/*
		Group Route
	*/

	group := route.Group("/api/v1")

	group.GET("/notes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
