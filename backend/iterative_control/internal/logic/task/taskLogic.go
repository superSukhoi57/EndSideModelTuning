package task

import (
	"context"
	"fmt"

	"iterative_control/internal/model"
	"iterative_control/internal/svc"
	"iterative_control/internal/types"
)

type TaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogic {
	return &TaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskLogic) Create(req *types.TaskCreateReq) error {
	task := &model.Task{
		ID:          req.ID,
		ParameterID: req.ParameterID,
		UserID:      req.UserID,
		Desc:        req.Desc,
	}

	if err := l.svcCtx.TaskDAO.Create(l.ctx, task); err != nil {
		return fmt.Errorf("create task failed: %w", err)
	}

	return nil
}

func (l *TaskLogic) Update(req *types.TaskUpdateReq) error {
	existing, err := l.svcCtx.TaskDAO.FindByID(l.ctx, req.ID)
	if err != nil {
		return fmt.Errorf("find task failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("task not found")
	}

	updates := &model.Task{ID: req.ID}
	if req.Desc != "" {
		updates.Desc = req.Desc
	}

	if err := l.svcCtx.TaskDAO.Update(l.ctx, updates); err != nil {
		return fmt.Errorf("update task failed: %w", err)
	}

	return nil
}

func (l *TaskLogic) Delete(id int64) error {
	existing, err := l.svcCtx.TaskDAO.FindByID(l.ctx, id)
	if err != nil {
		return fmt.Errorf("find task failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("task not found")
	}

	if err := l.svcCtx.TaskDAO.Delete(l.ctx, id); err != nil {
		return fmt.Errorf("delete task failed: %w", err)
	}

	return nil
}

func (l *TaskLogic) GetByID(id int64) (*types.TaskResp, error) {
	task, err := l.svcCtx.TaskDAO.FindByID(l.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find task failed: %w", err)
	}
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}

	return &types.TaskResp{
		ID:          task.ID,
		ParameterID: task.ParameterID,
		UserID:      task.UserID,
		Desc:        task.Desc,
		CreateAt:    task.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:    task.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (l *TaskLogic) List(req *types.TaskListReq) (*types.PageResp, error) {
	result, err := l.svcCtx.TaskDAO.List(l.ctx, req.Page, req.PageSize, req.UserID, req.ParameterID, req.Desc)
	if err != nil {
		return nil, fmt.Errorf("list tasks failed: %w", err)
	}

	list := make([]types.TaskResp, 0, len(result.List))
	for _, t := range result.List {
		list = append(list, types.TaskResp{
			ID:          t.ID,
			ParameterID: t.ParameterID,
			UserID:      t.UserID,
			Desc:        t.Desc,
			CreateAt:    t.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:    t.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.PageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
