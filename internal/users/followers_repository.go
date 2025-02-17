package users

import (
	"context"

	"github.com/huxleyberg/socialworks/internal/models"
	"github.com/huxleyberg/socialworks/pkg"
	"gorm.io/gorm"
)

type FollowerRepository struct {
	DB *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) FollowerRepository {
	return FollowerRepository{DB: db}
}

func (f *FollowerRepository) Follow(ctx context.Context, followerID, userID int64) error {
	follower := models.Follower{UserID: userID, FollowerID: followerID}
	if err := f.DB.WithContext(ctx).Create(&follower).Error; err != nil {
		if gorm.ErrDuplicatedKey == err {
			return pkg.ErrConflict
		}
		return err
	}
	return nil
}

func (f *FollowerRepository) Unfollow(ctx context.Context, followerID, userID int64) error {
	return f.DB.WithContext(ctx).Where("user_id = ? AND follower_id = ?", userID, followerID).Delete(&models.Follower{}).Error
}
