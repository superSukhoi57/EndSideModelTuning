package logic

import (
	"auth/internal/svc"
	"auth/internal/types"
	"context"
	"time"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (*types.LogoutResp, error) {
	if err := l.svcCtx.TokenDAO.DeleteAllTokens(l.ctx, req.UserID); err != nil {
		return nil, err
	}

	return &types.LogoutResp{
		Base: types.BaseResp{
			Code:      0,
			Message:   "success",
			Timestamp: time.Now().Unix(),
		},
	}, nil
}
