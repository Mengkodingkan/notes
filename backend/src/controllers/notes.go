package notes

import (
	"fmt"
	model "mengkodingkan/notes/src/database/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllNotes(db *gorm.DB) (*[]model.Note, string)
	GetNoteById(db *gorm.DB, id string) (*model.Note, string)
	CreateNote(db *gorm.DB, data *model.Note) (*model.Note, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllNotes() (*[]model.Note, string) {
	var notes []model.Note

	db := r.db.Find(&notes)
	err := make(chan string, 1)

	results := db.Debug().Select("*").Find(&notes)
	if results.Error != nil {
		err <- "404_NOT_FOUND"

		return &notes, <-err
	} else {
		err <- "nil"
	}

	return &notes, <-err
}

func (r *repository) GetNoteById(id string) (*model.Note, string) {
	var note model.Note

	db := r.db.Where("id = ?", id).First(&note)
	err := make(chan string, 1)

	results := db.Debug().Select("*").First(&note)

	if results.Error != nil {
		err <- "404_NOT_FOUND"

		return &note, <-err
	} else {
		err <- "nil"
	}

	return &note, <-err
}

func (r *repository) CreateNote(data *model.Note) (*model.Note, string) {
	var note model.Note

	fmt.Print(data)
	if data.Title == "" || data.Content == "" {
		return &note, "TITLE_OR_CONTENT_EMPTY"
	}

	db := r.db.Model(&note)
	err := make(chan string, 1)

	note.ID = uuid.New().String()
	note.Title = data.Title
	note.Content = data.Content
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	add := db.Debug().Create(&note)
	db.Commit()

	if add.Error != nil {
		err <- "FAILED_TO_CREATE_NOTE"
		return &note, <-err
	} else {
		err <- "nil"
	}

	return &note, <-err
}
