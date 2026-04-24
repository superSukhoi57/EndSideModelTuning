package logic

import (
	"auth/internal/svc"
	"auth/internal/types"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type StateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StateLogic {
	return &StateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StateLogic) State() (*types.StateResp, error) {
	state, err := l.svcCtx.StateDAO.GenerateAndStoreState(l.ctx, 5*time.Minute)
	if err != nil {
		logx.Errorf("Failed to generate state: %v", err)
		return nil, fmt.Errorf("failed to generate state")
	}

	return &types.StateResp{
		Base: types.BaseResp{
			Code:      0,
			Message:   "success",
			Timestamp: time.Now().Unix(),
		},
		State: state,
	}, nil
}
