package task

import (
	"context"
	"fmt"
	"log"

	"backend/common/enumeration"
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
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Create task - UserID: %d", userID)

	task := &model.Task{
		ID:          req.ID,
		ParameterID: req.ParameterID,
		UserID:      userID,
		Desc:        req.Desc,
	}

	if err := l.svcCtx.TaskDAO.Create(l.ctx, task); err != nil {
		return fmt.Errorf("create task failed: %w", err)
	}

	return nil
}

func (l *TaskLogic) Update(req *types.TaskUpdateReq) error {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Update task - UserID: %d, TaskID: %d", userID, req.ID)

	existing, err := l.svcCtx.TaskDAO.FindByID(l.ctx, req.ID)
	if err != nil {
		return fmt.Errorf("find task failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("task not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("task not owned by current user")
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
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return fmt.Errorf("missing user id in context")
	}
	log.Printf("Delete task - UserID: %d, TaskID: %d", userID, id)

	existing, err := l.svcCtx.TaskDAO.FindByID(l.ctx, id)
	if err != nil {
		return fmt.Errorf("find task failed: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("task not found")
	}
	if existing.UserID != userID {
		return fmt.Errorf("task not owned by current user")
	}

	if err := l.svcCtx.TaskDAO.Delete(l.ctx, id); err != nil {
		return fmt.Errorf("delete task failed: %w", err)
	}

	return nil
}

func (l *TaskLogic) GetByID(id int64) (*types.TaskResp, error) {
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("Get task by ID - UserID: %d, TaskID: %d", userID, id)

	task, err := l.svcCtx.TaskDAO.FindByID(l.ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find task failed: %w", err)
	}
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}
	if task.UserID != userID {
		return nil, fmt.Errorf("task not owned by current user")
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
	userID, ok := l.ctx.Value(enumeration.UserIDKey).(int64)
	if !ok {
		return nil, fmt.Errorf("missing user id in context")
	}
	log.Printf("List tasks - UserID: %d", userID)

	result, err := l.svcCtx.TaskDAO.List(l.ctx, req.Page, req.PageSize, userID, req.ParameterID, req.Desc)
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
