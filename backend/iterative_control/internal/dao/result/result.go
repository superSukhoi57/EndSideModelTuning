package result

import (
	"context"
	"fmt"

	"iterative_control/internal/model"

	"gorm.io/gorm"
)

type ResultDAO struct {
	db *gorm.DB
}

func NewResultDAO(db *gorm.DB) *ResultDAO {
	return &ResultDAO{db: db}
}

func (d *ResultDAO) Create(ctx context.Context, result *model.Result) error {
	if err := d.db.WithContext(ctx).Create(result).Error; err != nil {
		return fmt.Errorf("failed to create result: %w", err)
	}
	return nil
}

func (d *ResultDAO) Update(ctx context.Context, result *model.Result) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Result{}).
		Where("id = ? AND deleted = 0", result.ID).
		Updates(result).Error; err != nil {
		return fmt.Errorf("failed to update result: %w", err)
	}
	return nil
}

func (d *ResultDAO) Delete(ctx context.Context, id int64) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Result{}).
		Where("id = ?", id).
		Update("deleted", 1).Error; err != nil {
		return fmt.Errorf("failed to delete result: %w", err)
	}
	return nil
}

func (d *ResultDAO) FindByID(ctx context.Context, id int64) (*model.Result, error) {
	var result model.Result
	err := d.db.WithContext(ctx).
		Where("id = ? AND deleted = 0", id).
		First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find result by id: %w", err)
	}
	return &result, nil
}

type PageResult struct {
	Total int64
	List  []model.Result
}

func (d *ResultDAO) List(ctx context.Context, page, pageSize int, userID int64, machineID int64, desc string) (*PageResult, error) {
	var results []model.Result
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Result{}).Where("deleted = 0")

	if userID > 0 {
		query = query.Where("userid = ?", userID)
	}
	if machineID > 0 {
		query = query.Where("machineid = ?", machineID)
	}
	if desc != "" {
		query = query.Where("desc LIKE ?", "%"+desc+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count results: %w", err)
	}

	offset := (page - 1) * pageSize
	if err := query.Order("createAt DESC").Offset(offset).Limit(pageSize).Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to list results: %w", err)
	}

	return &PageResult{
		Total: total,
		List:  results,
	}, nil
}
