package posts

import (
	"context"

	"github.com/huxleyberg/socialworks/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return CommentRepository{DB: db}
}

func (c *CommentRepository) GetByPostID(ctx context.Context, postID int64) ([]models.Comment, error) {
	var comments []models.Comment
	result := c.DB.WithContext(ctx).
		Joins("User").
		Where("post_id = ?", postID).
		Order("created_at DESC").
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (c *CommentRepository) Create(ctx context.Context, comment *models.Comment) error {
	return c.DB.WithContext(ctx).Create(comment).Error
}
