package types

type CallbackReq struct {
	Code  string `form:"code"`
	State string `form:"state"`
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
	Base     BaseResp `json:"base"`
	UserInfo UserInfo `json:"user_info"`
}
