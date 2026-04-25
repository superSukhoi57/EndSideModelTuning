package machine

import (
	"context"
	"fmt"

	"iterative_control/internal/model"

	"gorm.io/gorm"
)

type MachineDAO struct {
	db *gorm.DB
}

func NewMachineDAO(db *gorm.DB) *MachineDAO {
	return &MachineDAO{db: db}
}

func (d *MachineDAO) Create(ctx context.Context, machine *model.Machine) error {
	if err := d.db.WithContext(ctx).Create(machine).Error; err != nil {
		return fmt.Errorf("failed to create machine: %w", err)
	}
	return nil
}

func (d *MachineDAO) Update(ctx context.Context, machine *model.Machine) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Machine{}).
		Where("id = ? AND deleted = 0", machine.ID).
		Updates(machine).Error; err != nil {
		return fmt.Errorf("failed to update machine: %w", err)
	}
	return nil
}

func (d *MachineDAO) Delete(ctx context.Context, id int64) error {
	if err := d.db.WithContext(ctx).
		Model(&model.Machine{}).
		Where("id = ?", id).
		Update("deleted", 1).Error; err != nil {
		return fmt.Errorf("failed to delete machine: %w", err)
	}
	return nil
}

func (d *MachineDAO) FindByID(ctx context.Context, id int64) (*model.Machine, error) {
	var machine model.Machine
	err := d.db.WithContext(ctx).
		Where("id = ? AND deleted = 0", id).
		First(&machine).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find machine by id: %w", err)
	}
	return &machine, nil
}

type PageResult struct {
	Total int64
	List  []model.Machine
}

func (d *MachineDAO) List(ctx context.Context, page, pageSize int, ip string, userID int64, isFinsh *int8) (*PageResult, error) {
	var machines []model.Machine
	var total int64

	query := d.db.WithContext(ctx).Model(&model.Machine{}).Where("deleted = 0")

	if ip != "" {
		query = query.Where("ip LIKE ?", "%"+ip+"%")
	}
	if userID > 0 {
		query = query.Where("userid = ?", userID)
	}
	if isFinsh != nil {
		query = query.Where("isfinsh = ?", *isFinsh)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count machines: %w", err)
	}

	offset := (page - 1) * pageSize
	if err := query.Order("createAt DESC").Offset(offset).Limit(pageSize).Find(&machines).Error; err != nil {
		return nil, fmt.Errorf("failed to list machines: %w", err)
	}

	return &PageResult{
		Total: total,
		List:  machines,
	}, nil
}
