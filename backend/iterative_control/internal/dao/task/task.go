package task

import (
	"context"
	"fmt"

	"iterative_control/internal/model"

	"gorm.io/gorm"
)

type TaskDAO struct {
	db *gorm.DB
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{db: db}
}

func (d *TaskDAO) Create(ctx context.Context, task *model.Task) error {
	if err := d.db.WithContext(ctx).Create(task).Error; err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (d *TaskDAO) Update(ctx context.Context, task *model.Task) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ? AND deleted = 0", task.ID).
		Updates(task).Error; err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (d *TaskDAO) Delete(ctx context.Context, id int64) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ?", id).
		Update("deleted", 1).Error; err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (d *TaskDAO) FindByID(ctx context.Context, id int64) (*model.Task, error) {
	var task model.Task
	err := d.db.WithContext(ctx).
		Where("id = ? AND deleted = 0", id).
		First(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find task by id: %w", err)
	}
	return &task, nil
}

type PageResult struct {
	Total int64
	List  []model.Task
}

func (d *TaskDAO) List(ctx context.Context, page, pageSize int, userID int64, parameterID int64, desc string) (*PageResult, error) {
	var tasks []model.Task
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Task{}).Where("deleted = 0")

	if userID > 0 {
		query = query.Where("userid = ?", userID)
	}
	if parameterID > 0 {
		query = query.Where("paramterid = ?", parameterID)
	}
	if desc != "" {
		query = query.Where("desc LIKE ?", "%"+desc+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count tasks: %w", err)
	}

	offset := (page - 1) * pageSize
	if err := query.Order("createAt DESC").Offset(offset).Limit(pageSize).Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	return &PageResult{
		Total: total,
		List:  tasks,
	}, nil
}
