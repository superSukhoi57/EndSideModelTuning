package types

import "errors"

var (
	ErrTokenRevoked = errors.New("token has been revoked")
)

type CallbackReq struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type StateReq struct{}

type StateResp struct {
	Base  BaseResp `json:"base"`
	State string   `json:"state"`
}

type BaseResp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type UserInfo struct {
	OpenId     string `json:"open_id"`
	UnionId    string `json:"union_id"`
	Name       string `json:"name"`
	EnName     string `json:"en_name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	AvatarUrl  string `json:"avatar_url"`
	TenantName string `json:"tenant_name"`
}

type CallbackResp struct {
	Base         BaseResp `json:"base"`
	UserInfo     UserInfo `json:"user_info"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	Base         BaseResp `json:"base"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
}

type VerifyTokenReq struct {
	AccessToken string `form:"access_token"`
}

type VerifyTokenResp struct {
	Base   BaseResp `json:"base"`
	UserID int64    `json:"user_id"`
	Role   string   `json:"role"`
}

type LogoutReq struct {
	UserID int64 `json:"-"`
}

type LogoutResp struct {
	Base BaseResp `json:"base"`
}
