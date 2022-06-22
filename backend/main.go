package main

import (
	"log" // log package is used to log messages
	route "mengkodingkan/notes/src/routes"
	util "mengkodingkan/notes/src/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := _init()

	log.Fatal(router.Run(":" + util.Get("PORT")))
}

func _init() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.InitNotesRoutes(r)

	return r
}
