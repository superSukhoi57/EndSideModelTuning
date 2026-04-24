package logic

import (
	"auth/internal/svc"
	"auth/internal/types"
	"context"
	"strconv"
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
	//直接在这里拦截
	claims, err := l.svcCtx.JWTMgr.ValidateToken(req.AccessToken, false)
	if err != nil {
		return &types.LogoutResp{
			Base: types.BaseResp{
				Code:      0,
				Message:   "success",
				Timestamp: time.Now().Unix(),
			},
		}, nil
	}

	userID, err := strconv.ParseInt(claims.UserID, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.TokenDAO.DeleteAllTokens(l.ctx, userID); err != nil {
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
