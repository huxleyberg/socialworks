package users

import (
	"context"

	"github.com/huxleyberg/socialworks/internal/models"
	"github.com/huxleyberg/socialworks/pkg"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (u *UserRepository) Create(ctx context.Context, user *models.User) error {
	return u.DB.WithContext(ctx).Create(user).Error
}

func (u *UserRepository) GetByID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	result := u.DB.WithContext(ctx).First(&user, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, pkg.ErrNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}
