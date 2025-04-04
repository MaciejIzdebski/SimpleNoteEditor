package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint `gorm:"primarykey"`
	Title      string 
	Description string 
	Owner string 	  
    CreatedAt time.Time 
    UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

