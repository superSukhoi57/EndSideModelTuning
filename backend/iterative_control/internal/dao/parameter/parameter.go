package parameter

import (
	"context"
	"fmt"

	"iterative_control/internal/model"

	"gorm.io/gorm"
)

type ParameterDAO struct {
	db *gorm.DB
}

func NewParameterDAO(db *gorm.DB) *ParameterDAO {
	return &ParameterDAO{db: db}
}

func (d *ParameterDAO) Create(ctx context.Context, param *model.Parameter) error {
	if err := d.db.WithContext(ctx).Create(param).Error; err != nil {
		return fmt.Errorf("failed to create parameter: %w", err)
	}
	return nil
}

func (d *ParameterDAO) Update(ctx context.Context, param *model.Parameter) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Parameter{}).
		Where("id = ? AND deleted = 0", param.ID).
		Updates(param).Error; err != nil {
		return fmt.Errorf("failed to update parameter: %w", err)
	}
	return nil
}

func (d *ParameterDAO) Delete(ctx context.Context, id int64) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Parameter{}).
		Where("id = ?", id).
		Update("deleted", 1).Error; err != nil {
		return fmt.Errorf("failed to delete parameter: %w", err)
	}
	return nil
}

func (d *ParameterDAO) FindByID(ctx context.Context, id int64) (*model.Parameter, error) {
	var param model.Parameter
	err := d.db.WithContext(ctx).
		Where("id = ? AND deleted = 0", id).
		First(&param).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find parameter by id: %w", err)
	}
	return &param, nil
}

type PageResult struct {
	Total int64
	List  []model.Parameter
}

func (d *ParameterDAO) List(ctx context.Context, page, pageSize int, userID int64, desc string) (*PageResult, error) {
	var params []model.Parameter
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Parameter{}).Where("deleted = 0")

	if userID > 0 {
		query = query.Where("userid = ?", userID)
	}
	if desc != "" {
		query = query.Where("desc LIKE ?", "%"+desc+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count parameters: %w", err)
	}

	offset := (page - 1) * pageSize
	if err := query.Order("createAt DESC").Offset(offset).Limit(pageSize).Find(&params).Error; err != nil {
		return nil, fmt.Errorf("failed to list parameters: %w", err)
	}

	return &PageResult{
		Total: total,
		List:  params,
	}, nil
}
