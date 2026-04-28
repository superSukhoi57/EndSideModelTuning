package parameter

import (
	"context"
	"fmt"
	"log"

	"backend/common/enumeration"
	"iterative_control/internal/model"
	"iterative_control/internal/svc"
	"iterative_control/internal/types"
)

type ParameterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewParameterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParameterLogic {
	return &ParameterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ParameterLogic) Create(req *types.ParameterCreateReq) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Create parameter - UserID: %d", userID)

	param := &model.Parameter{
		ID:         req.ID,
		UserID:     userID,
		Parameters: req.Parameters,
		Script:     req.Script,
		Desc:       req.Desc,
	}

	if err := l.svcCtx.ParameterDAO.Create(l.ctx, param); err != nil {
		return fmt.Errorf("create parameter failed: %w", err)
	}

	return nil
}

func (l *ParameterLogic) Update(req *types.ParameterUpdateReq) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Update parameter - UserID: %d, ParameterID: %d", userID, req.ID)

	existing, err := l.svcCtx.ParameterDAO.FindByID(l.ctx, req.ID)
	if err != nil {
		return fmt.Errorf("find parameter failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("parameter not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("parameter not owned by current user")
	}

	updates := &model.Parameter{ID: req.ID}
	if req.Parameters != "" {
		updates.Parameters = req.Parameters
	}
	if req.Script != "" {
		updates.Script = req.Script
	}
	if req.Desc != "" {
		updates.Desc = req.Desc
	}

	if err := l.svcCtx.ParameterDAO.Update(l.ctx, updates); err != nil {
		return fmt.Errorf("update parameter failed: %w", err)
	}

	return nil
}

func (l *ParameterLogic) Delete(id int64) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Delete parameter - UserID: %d, ParameterID: %d", userID, id)

	existing, err := l.svcCtx.ParameterDAO.FindByID(l.ctx, id)
	if err != nil {
		return fmt.Errorf("find parameter failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("parameter not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("parameter not owned by current user")
	}

	if err := l.svcCtx.ParameterDAO.Delete(l.ctx, id); err != nil {
		return fmt.Errorf("delete parameter failed: %w", err)
	}

	return nil
}

func (l *ParameterLogic) GetByID(id int64) (*types.ParameterResp, error) {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("Get parameter by ID - UserID: %d, ParameterID: %d", userID, id)

	param, err := l.svcCtx.ParameterDAO.FindByID(l.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find parameter failed: %w", err)
	}
	if param == nil {
		return nil, fmt.Errorf("parameter not found")
	}
	if param.UserID != userID {
		return nil, fmt.Errorf("parameter not owned by current user")
	}

	return &types.ParameterResp{
		ID:         param.ID,
		UserID:     param.UserID,
		Parameters: param.Parameters,
		Script:     param.Script,
		Desc:       param.Desc,
		CreateAt:   param.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:   param.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (l *ParameterLogic) List(req *types.ParameterListReq) (*types.PageResp, error) {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("List parameters - UserID: %d", userID)

	result, err := l.svcCtx.ParameterDAO.List(l.ctx, req.Page, req.PageSize, userID, req.Desc)
	if err != nil {
		return nil, fmt.Errorf("list parameters failed: %w", err)
	}

	list := make([]types.ParameterResp, 0, len(result.List))
	for _, p := range result.List {
		list = append(list, types.ParameterResp{
			ID:         p.ID,
			UserID:     p.UserID,
			Parameters: p.Parameters,
			Script:     p.Script,
			Desc:       p.Desc,
			CreateAt:   p.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:   p.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.PageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
