package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	model "mengkodingkan/notes/src/models"
)

func Connection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(os.Getenv("test.db")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database!")
	}

	err = db.AutoMigrate(&model.Note{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
