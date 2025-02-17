package models

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   string         `json:"content"`
	Title     string         `json:"title"`
	UserID    int64          `json:"user_id"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Version   int            `json:"version"`
	Comments  []Comment      `json:"comments"`
	User      User           `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentsCount int `json:"comment_count"`
}
