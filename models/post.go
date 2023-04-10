package models

import (
	"time"
)

type Post struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title" gorm:"not null"`
	Body      string     `json:"body" gorm:"not null"`
	UserID    uint       `json:"userId"`
	CreatedBy string     `json:"createdBy"`
	CreatedAt *time.Time `json:"created_at"`
}
