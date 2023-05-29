package models

import (
	"time"
)

type Like struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	PostID    uint      `json:"postId"`
	CreatedBy uint      `json:"createdBy"`
	CreatedAt time.Time `json:"created_at"`
}
