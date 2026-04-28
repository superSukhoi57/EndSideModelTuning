package result

import (
	"context"
	"fmt"
	"log"

	"backend/common/enumeration"
	"iterative_control/internal/model"
	"iterative_control/internal/svc"
	"iterative_control/internal/types"
)

type ResultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResultLogic {
	return &ResultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResultLogic) Create(req *types.ResultCreateReq) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Create result - UserID: %d", userID)

	result := &model.Result{
		Result:    req.Result,
		UserID:    userID,
		MachineID: req.MachineID,
		Desc:      req.Desc,
	}

	if err := l.svcCtx.ResultDAO.Create(l.ctx, result); err != nil {
		return fmt.Errorf("create result failed: %w", err)
	}

	return nil
}

func (l *ResultLogic) Update(req *types.ResultUpdateReq) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Update result - UserID: %d, ResultID: %d", userID, req.ID)

	existing, err := l.svcCtx.ResultDAO.FindByID(l.ctx, req.ID)
	if err != nil {
		return fmt.Errorf("find result failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("result not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("result not owned by current user")
	}

	updates := &model.Result{ID: req.ID}
	if req.Result != "" {
		updates.Result = req.Result
	}
	if req.Desc != "" {
		updates.Desc = req.Desc
	}

	if err := l.svcCtx.ResultDAO.Update(l.ctx, updates); err != nil {
		return fmt.Errorf("update result failed: %w", err)
	}

	return nil
}

func (l *ResultLogic) Delete(id int64) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Delete result - UserID: %d, ResultID: %d", userID, id)

	existing, err := l.svcCtx.ResultDAO.FindByID(l.ctx, id)
	if err != nil {
		return fmt.Errorf("find result failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("result not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("result not owned by current user")
	}

	if err := l.svcCtx.ResultDAO.Delete(l.ctx, id); err != nil {
		return fmt.Errorf("delete result failed: %w", err)
	}

	return nil
}

func (l *ResultLogic) GetByID(id int64) (*types.ResultResp, error) {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("Get result by ID - UserID: %d, ResultID: %d", userID, id)

	result, err := l.svcCtx.ResultDAO.FindByID(l.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find result failed: %w", err)
	}
	if result == nil {
		return nil, fmt.Errorf("result not found")
	}
	if result.UserID != userID {
		return nil, fmt.Errorf("result not owned by current user")
	}

	return &types.ResultResp{
		ID:        result.ID,
		Result:    result.Result,
		UserID:    result.UserID,
		MachineID: result.MachineID,
		Desc:      result.Desc,
		CreateAt:  result.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  result.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (l *ResultLogic) List(req *types.ResultListReq) (*types.PageResp, error) {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("List results - UserID: %d", userID)

	result, err := l.svcCtx.ResultDAO.List(l.ctx, req.Page, req.PageSize, userID, req.MachineID, req.Desc)
	if err != nil {
		return nil, fmt.Errorf("list results failed: %w", err)
	}

	list := make([]types.ResultResp, 0, len(result.List))
	for _, r := range result.List {
		list = append(list, types.ResultResp{
			ID:        r.ID,
			Result:    r.Result,
			UserID:    r.UserID,
			MachineID: r.MachineID,
			Desc:      r.Desc,
			CreateAt:  r.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:  r.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.PageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
