package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint `gorm:"primarykey" json:"id"`
	Title      string `json:"title"`
	Descrition string `json:"description"`
	Owner string 	  `json:"owner"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

