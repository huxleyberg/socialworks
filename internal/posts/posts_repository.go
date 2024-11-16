package posts

import (
	"context"

	"github.com/huxleyberg/socialworks/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return PostRepository{DB: db}
}

func (p *PostRepository) Create(ctx context.Context, post *models.Post) error {
	return p.DB.WithContext(ctx).Save(post).Error
}
