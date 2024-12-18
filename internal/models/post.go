package models

import (
	"github.com/lib/pq"
)

type Post struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   string         `json:"content"`
	Title     string         `json:"title"`
	UserID    int64          `json:"user_id"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}
