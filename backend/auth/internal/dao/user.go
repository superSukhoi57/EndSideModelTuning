package dao

import (
	"context"
	"fmt"

	"auth/internal/model"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d *UserDAO) FindByUnionID(ctx context.Context, unionID string) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).
		Where("lark_union_id = ? AND deleted = 0", unionID).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by union_id: %w", err)
	}
	return &user, nil
}

func (d *UserDAO) Create(ctx context.Context, user *model.User) error {
	if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (d *UserDAO) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).
		Where("id = ? AND deleted = 0", id).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}
	return &user, nil
}
