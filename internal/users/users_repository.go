package users

import (
	"context"
	"errors"

	"github.com/huxleyberg/socialworks/internal/models"
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
			return nil, errors.New("record not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
