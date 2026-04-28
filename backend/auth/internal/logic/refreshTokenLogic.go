package logic

import (
	"auth/internal/svc"
	"auth/internal/types"
	"context"
	"strconv"
	"time"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
刷新token
一般是在access_token过期后，刷新access_token和refresh_token
返回新的access_token和refresh_token
如果refresh_token过期，返回错误
如果refresh_token不存在redis，返回错误
*/
func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq, userID int64) (*types.RefreshTokenResp, error) {
	claims, err := l.svcCtx.JWTMgr.ValidateToken(req.RefreshToken, true)
	if err != nil {
		return nil, err
	}

	claimedUserID, err := strconv.ParseInt(claims.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	if claimedUserID != userID {
		return nil, types.ErrTokenRevoked
	}

	savedToken, err := l.svcCtx.TokenDAO.GetRefreshToken(l.ctx, userID)
	if err != nil {
		return nil, err
	}
	if savedToken == "" {
		return nil, types.ErrTokenRevoked
	}
	if savedToken != req.RefreshToken {
		return nil, types.ErrTokenRevoked
	}

	// 刷新token，这里继承过期时间
	tokenPair, err := l.svcCtx.JWTMgr.RefreshToken(req.RefreshToken, true)
	if err != nil {
		return nil, err
	}

	if err := l.svcCtx.TokenDAO.SaveAccessToken(l.ctx, userID, tokenPair.AccessToken, l.svcCtx.Config.JWT.AccessExpire); err != nil {
		return nil, err
	}
	if err := l.svcCtx.TokenDAO.SaveRefreshToken(l.ctx, userID, tokenPair.RefreshToken, l.svcCtx.Config.JWT.RefreshExpire); err != nil {
		return nil, err
	}

	return &types.RefreshTokenResp{
		Base: types.BaseResp{
			Code:      0,
			Message:   "success",
			Timestamp: time.Now().Unix(),
		},
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}
