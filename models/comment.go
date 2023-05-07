package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Comment   string    `json:"comment" gorm:"not null"`
	PostID    uint      `json:"postId"`
	CreatedBy uint      `json:"createdBy"`
	CreatedAt time.Time `json:"created_at"`
}
