package main

import (
	"log" // log package is used to log messages
	database "mengkodingkan/notes/src/database"
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

	db := database.Connection()

	route.InitNotesRoutes(db, r)

	return r
}
