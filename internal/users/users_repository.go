package users

import (
	"context"

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
	return u.DB.WithContext(ctx).Save(user).Error
}
