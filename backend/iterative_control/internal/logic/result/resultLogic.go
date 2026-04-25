package result

import (
	"context"
	"fmt"

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
	result := &model.Result{
		Result:    req.Result,
		UserID:    req.UserID,
		MachineID: req.MachineID,
		Desc:      req.Desc,
	}

	if err := l.svcCtx.ResultDAO.Create(l.ctx, result); err != nil {
		return fmt.Errorf("create result failed: %w", err)
	}

	return nil
}

func (l *ResultLogic) Update(req *types.ResultUpdateReq) error {
	existing, err := l.svcCtx.ResultDAO.FindByID(l.ctx, req.ID)
	if err != nil {
		return fmt.Errorf("find result failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("result not found")
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
	existing, err := l.svcCtx.ResultDAO.FindByID(l.ctx, id)
	if err != nil {
		return fmt.Errorf("find result failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("result not found")
	}

	if err := l.svcCtx.ResultDAO.Delete(l.ctx, id); err != nil {
		return fmt.Errorf("delete result failed: %w", err)
	}

	return nil
}

func (l *ResultLogic) GetByID(id int64) (*types.ResultResp, error) {
	result, err := l.svcCtx.ResultDAO.FindByID(l.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find result failed: %w", err)
	}
	if result == nil {
		return nil, fmt.Errorf("result not found")
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
	result, err := l.svcCtx.ResultDAO.List(l.ctx, req.Page, req.PageSize, req.UserID, req.MachineID, req.Desc)
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
