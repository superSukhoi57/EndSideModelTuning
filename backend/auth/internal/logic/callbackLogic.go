package logic

import (
	"auth/internal/model"
	"auth/internal/svc"
	"auth/internal/types"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	commonjwt "backend/common/jwt"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallbackLogic) Callback(req *types.CallbackReq) (*types.CallbackResp, error) {

	logx.Infof("Callback req: %v", req)
	if req.State == "" {
		return nil, errors.New("missing state")
	}
	valid, err := l.svcCtx.StateDAO.ValidateState(l.ctx, req.State)
	if err != nil {
		return nil, fmt.Errorf("failed to validate state: %w", err)
	}
	if !valid {
		return nil, errors.New("invalid state")
	}
	if req.Code == "" {
		return nil, errors.New("missing authorization code")
	}

	tokenResp, err := l.getAccessToken(req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	userInfo, err := l.getUserInfo(tokenResp.Data.AccessToken, tokenResp.Data.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	user, err := l.svcCtx.UserDAO.FindByUnionID(l.ctx, userInfo.UnionId)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if user == nil {
		node, _ := snowflake.NewNode(1)
		user = &model.User{
			ID:          node.Generate().Int64(),
			LarkUnionID: userInfo.UnionId,
			LarkAvatar:  userInfo.AvatarUrl,
			Username:    userInfo.Name,
			Role:        `{"role":"user"}`,
			Deleted:     0,
		}
		if err := l.svcCtx.UserDAO.Create(l.ctx, user); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	claims := commonjwt.CustomClaims{
		UserID:   fmt.Sprintf("%d", user.ID),
		Role:     user.Role,
		Username: user.Username,
	}

	tokenPair, err := l.svcCtx.JWTMgr.GenerateTokenPair(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	if err := l.svcCtx.TokenDAO.SaveAccessToken(l.ctx, user.ID, tokenPair.AccessToken, l.svcCtx.Config.JWT.AccessExpire); err != nil {
		return nil, fmt.Errorf("failed to save access token: %w", err)
	}
	if err := l.svcCtx.TokenDAO.SaveRefreshToken(l.ctx, user.ID, tokenPair.RefreshToken, l.svcCtx.Config.JWT.RefreshExpire); err != nil {
		return nil, fmt.Errorf("failed to save refresh token: %w", err)
	}

	resp := &types.CallbackResp{
		Base: types.BaseResp{
			Code:      0,
			Message:   "success",
			Timestamp: time.Now().Unix(),
		},
		UserInfo:     *userInfo,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}

	return resp, nil
}

type AppAccessTokenResp struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

type UserAccessTokenResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken      string `json:"access_token"`
		RefreshToken     string `json:"refresh_token"`
		ExpiresIn        int    `json:"expires_in"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		TokenType        string `json:"token_type"`
		OpenId           string `json:"open_id"`
		Name             string `json:"name"`
		EnName           string `json:"en_name"`
		AvatarUrl        string `json:"avatar_url"`
		Email            string `json:"email"`
		EnterpriseEmail  string `json:"enterprise_email"`
		Mobile           string `json:"mobile"`
		TenantKey        string `json:"tenant_key"`
		UnionId          string `json:"union_id"`
	} `json:"data"`
}

func (l *CallbackLogic) getAppAccessToken() (string, error) {
	tokenURL := "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"

	reqBody := map[string]string{
		"app_id":     l.svcCtx.Config.FeishuAuth.AppID,
		"app_secret": l.svcCtx.Config.FeishuAuth.AppSecret,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(l.ctx, "POST", tokenURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var tokenResp AppAccessTokenResp
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if tokenResp.Code != 0 {
		return "", fmt.Errorf("feishu API error: code=%d, msg=%s", tokenResp.Code, tokenResp.Msg)
	}

	return tokenResp.AppAccessToken, nil
}

func (l *CallbackLogic) getAccessToken(code string) (*UserAccessTokenResp, error) {
	appAccessToken, err := l.getAppAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get app access token: %w", err)
	}

	tokenURL := "https://open.feishu.cn/open-apis/authen/v1/oidc/access_token"

	reqBody := map[string]string{
		"grant_type": "authorization_code",
		"code":       code,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(l.ctx, "POST", tokenURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+appAccessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var tokenResp UserAccessTokenResp
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if tokenResp.Code != 0 {
		return nil, fmt.Errorf("feishu API error: code=%d, msg=%s", tokenResp.Code, tokenResp.Msg)
	}

	return &tokenResp, nil
}

type UserInfoResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name      string `json:"name"`
		EnName    string `json:"en_name"`
		Email     string `json:"email"`
		Mobile    string `json:"mobile"`
		OpenId    string `json:"open_id"`
		UnionId   string `json:"union_id"`
		AvatarUrl string `json:"avatar_url"`
		TenantKey string `json:"tenant_key"`
	} `json:"data"`
}

func (l *CallbackLogic) getUserInfo(accessToken string, refreshToken string) (*types.UserInfo, error) {
	userInfoURL := "https://open.feishu.cn/open-apis/authen/v1/user_info"

	req, err := http.NewRequestWithContext(l.ctx, "GET", userInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var userInfoResp UserInfoResp
	if err := json.Unmarshal(body, &userInfoResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if userInfoResp.Code != 0 {
		return nil, fmt.Errorf("feishu API error: code=%d, msg=%s", userInfoResp.Code, userInfoResp.Msg)
	}

	userInfo := &types.UserInfo{
		OpenId:     userInfoResp.Data.OpenId,
		UnionId:    userInfoResp.Data.UnionId,
		Name:       userInfoResp.Data.Name,
		EnName:     userInfoResp.Data.EnName,
		Email:      userInfoResp.Data.Email,
		Mobile:     userInfoResp.Data.Mobile,
		AvatarUrl:  userInfoResp.Data.AvatarUrl,
		TenantName: userInfoResp.Data.TenantKey,
	}

	return userInfo, nil
}
