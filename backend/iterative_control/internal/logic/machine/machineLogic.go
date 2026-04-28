package machine

import (
	"context"
	"fmt"
	"log"

	"backend/common/enumeration"
	"iterative_control/internal/model"
	"iterative_control/internal/svc"
	"iterative_control/internal/types"
)

type MachineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMachineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MachineLogic {
	return &MachineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MachineLogic) Create(req *types.MachineCreateReq) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Create machine - UserID: %d", userID)

	machine := &model.Machine{
		IP:     req.IP,
		Pwd:    req.Pwd,
		UserID: userID,
		Core:   req.Core,
		RAM:    req.RAM,
		Memory: req.Memory,
		OS:     req.OS,
		Desc:   req.Desc,
	}

	if err := l.svcCtx.MachineDAO.Create(l.ctx, machine); err != nil {
		return fmt.Errorf("create machine failed: %w", err)
	}

	return nil
}

func (l *MachineLogic) Update(req *types.MachineUpdateReq) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Update machine - UserID: %d, MachineID: %d", userID, req.ID)

	existing, err := l.svcCtx.MachineDAO.FindByID(l.ctx, req.ID)
	if err != nil {
		return fmt.Errorf("find machine failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("machine not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("machine not owned by current user")
	}

	updates := &model.Machine{ID: req.ID}
	if req.IP != "" {
		updates.IP = req.IP
	}
	if req.Pwd != "" {
		updates.Pwd = req.Pwd
	}
	if req.IsFinsh != nil {
		updates.IsFinsh = *req.IsFinsh
	}
	if req.ResultID != nil {
		updates.ResultID = req.ResultID
	}
	if req.Core != nil {
		updates.Core = req.Core
	}
	if req.RAM != nil {
		updates.RAM = req.RAM
	}
	if req.Memory != nil {
		updates.Memory = req.Memory
	}
	if req.OS != "" {
		updates.OS = req.OS
	}
	if req.Desc != "" {
		updates.Desc = req.Desc
	}

	if err := l.svcCtx.MachineDAO.Update(l.ctx, updates); err != nil {
		return fmt.Errorf("update machine failed: %w", err)
	}

	return nil
}

func (l *MachineLogic) Delete(id int64) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Delete machine - UserID: %d, MachineID: %d", userID, id)

	existing, err := l.svcCtx.MachineDAO.FindByID(l.ctx, id)
	if err != nil {
		return fmt.Errorf("find machine failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("machine not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("machine not owned by current user")
	}

	if err := l.svcCtx.MachineDAO.Delete(l.ctx, id); err != nil {
		return fmt.Errorf("delete machine failed: %w", err)
	}

	return nil
}

func (l *MachineLogic) GetByID(id int64) (*types.MachineResp, error) {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("Get machine by ID - UserID: %d, MachineID: %d", userID, id)

	machine, err := l.svcCtx.MachineDAO.FindByID(l.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find machine failed: %w", err)
	}
	if machine == nil {
		return nil, fmt.Errorf("machine not found")
	}
	if machine.UserID != userID {
		return nil, fmt.Errorf("machine not owned by current user")
	}

	return &types.MachineResp{
		ID:       machine.ID,
		IP:       machine.IP,
		UserID:   machine.UserID,
		IsFinsh:  machine.IsFinsh,
		ResultID: machine.ResultID,
		Core:     machine.Core,
		RAM:      machine.RAM,
		Memory:   machine.Memory,
		OS:       machine.OS,
		Desc:     machine.Desc,
		CreateAt: machine.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: machine.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (l *MachineLogic) List(req *types.MachineListReq) (*types.PageResp, error) {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("List machines - UserID: %d", userID)

	result, err := l.svcCtx.MachineDAO.List(l.ctx, req.Page, req.PageSize, req.IP, userID, req.IsFinsh)
	if err != nil {
		return nil, fmt.Errorf("list machines failed: %w", err)
	}

	list := make([]types.MachineResp, 0, len(result.List))
	for _, m := range result.List {
		list = append(list, types.MachineResp{
			ID:       m.ID,
			IP:       m.IP,
			UserID:   m.UserID,
			IsFinsh:  m.IsFinsh,
			ResultID: m.ResultID,
			Core:     m.Core,
			RAM:      m.RAM,
			Memory:   m.Memory,
			OS:       m.OS,
			Desc:     m.Desc,
			CreateAt: m.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt: m.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.PageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
