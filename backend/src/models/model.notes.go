package notes

import "time"

type Note struct {
	ID        string    `gorm:"primary_key"`
	Title     string    `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
