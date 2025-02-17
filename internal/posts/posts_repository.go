package posts

import (
	"context"

	"github.com/huxleyberg/socialworks/internal/models"
	"github.com/huxleyberg/socialworks/internal/utils"
	"github.com/huxleyberg/socialworks/pkg"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return PostRepository{DB: db}
}

func (p *PostRepository) Create(ctx context.Context, post *models.Post) error {
	return p.DB.WithContext(ctx).Create(post).Error
}

func (p *PostRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	var post models.Post
	result := p.DB.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, pkg.ErrNotFound
		}
		return nil, result.Error
	}
	return &post, nil
}

func (p *PostRepository) Delete(ctx context.Context, postID int64) error {
	result := p.DB.WithContext(ctx).Delete(&models.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return pkg.ErrNotFound
	}
	return nil
}

func (p *PostRepository) Update(ctx context.Context, post *models.Post) error {
	result := p.DB.WithContext(ctx).Model(post).Where("id = ? AND version = ?", post.ID, post.Version).Updates(models.Post{
		Title:   post.Title,
		Content: post.Content,
		Version: post.Version + 1,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return pkg.ErrNotFound
	}
	return nil
}

func (p *PostRepository) GetUserFeed(ctx context.Context, userID int64, fq utils.PaginatedFeedQuery) ([]models.PostWithMetadata, error) {
	var feed []models.PostWithMetadata
	query := p.DB.WithContext(ctx).Table("posts p").
		Select("p.*, u.username, COUNT(c.id) AS comments_count").
		Joins("LEFT JOIN comments c ON c.post_id = p.id").
		Joins("LEFT JOIN users u ON p.user_id = u.id").
		Joins("JOIN followers f ON f.follower_id = p.user_id OR p.user_id = ?", userID).
		Where("f.user_id = ? AND (p.title ILIKE ? OR p.content ILIKE ?) AND (p.tags @> ? OR ? = '{}')", userID, "%"+fq.Search+"%", "%"+fq.Search+"%", fq.Tags, fq.Tags).
		Group("p.id, u.username").
		Order("p.created_at " + fq.Sort).
		Limit(fq.Limit).
		Offset(fq.Offset)

	if err := query.Scan(&feed).Error; err != nil {
		return nil, err
	}
	return feed, nil
}
