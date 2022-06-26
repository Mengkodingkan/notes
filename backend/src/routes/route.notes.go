package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	note "mengkodingkan/notes/src/controllers"
	model "mengkodingkan/notes/src/database/models"
	middleware "mengkodingkan/notes/src/middleware"
)

func InitNotesRoutes(db *gorm.DB, route *gin.Engine) {
	/*
		Controllers
	*/
	note := note.NewRepository(db)
	/*
		Group Route
	*/
	group := route.Group("/api/v1").Use(middleware.Auth())
	group.GET("/notes", func(c *gin.Context) {
		notes, err := note.GetAllNotes()

		if err == "404_NOT_FOUND" {
			c.JSON(404, gin.H{
				"message": "404_NOT_FOUND",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "success",
				"data":   notes,
			})
		}
	})

	group.GET("/notes/:id", func(c *gin.Context) {
		id := c.Param("id")
		note, err := note.GetNoteById(id)

		if err == "404_NOT_FOUND" {
			c.JSON(404, gin.H{
				"status":  "error",
				"message": "404_NOT_FOUND",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "success",
				"data":   note,
			})
		}
	})

	group.POST("/notes", func(c *gin.Context) {
		var noted model.Note
		c.ShouldBindJSON(&noted)
		create, err := note.CreateNote(&noted)

		switch err {
		case "nil":
			c.JSON(200, gin.H{
				"status": "success",
				"data":   create,
			})

		case "FAILED_TO_CREATE_NOTE":
			c.JSON(400, gin.H{
				"status":  "error",
				"message": "FAILED_TO_CREATE_NOTE",
			})

		case "TITLE_OR_CONTENT_EMPTY":
			c.JSON(400, gin.H{
				"status":  "error",
				"message": "TITLE_OR_CONTENT_EMPTY",
			})
		}
	})

	group.PUT("/notes/:id", func(c *gin.Context) {})
}
