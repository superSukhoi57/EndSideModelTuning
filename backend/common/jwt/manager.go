package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type Manager struct {
	accessSecret  []byte
	refreshSecret []byte
	accessExpire  time.Duration
	refreshExpire time.Duration
	issuer        string
}

func NewManager(cfg Config) *Manager {
	return &Manager{
		accessSecret:  []byte(cfg.AccessSecret),
		refreshSecret: []byte(cfg.RefreshSecret),
		accessExpire:  cfg.AccessExpire,
		refreshExpire: cfg.RefreshExpire,
		issuer:        cfg.Issuer,
	}
}

func (m *Manager) GenerateTokenPair(claims CustomClaims) (*TokenPair, error) {
	accessToken, err := m.generateToken(&claims, m.accessSecret, m.accessExpire)
	if err != nil {
		return nil, err
	}

	refreshToken, err := m.generateToken(&claims, m.refreshSecret, m.refreshExpire)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *Manager) RefreshToken(refreshTokenStr string, inheritExpiry bool) (*TokenPair, error) {
	claims, err := m.ValidateToken(refreshTokenStr, true)
	if err != nil {
		return nil, err
	}

	customClaims := CustomClaims{
		BaseClaims: BaseClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer: m.issuer,
			},
		},
		UserID:   claims.UserID,
		Role:     claims.Role,
		Username: claims.Username,
	}

	accessToken, err := m.generateToken(&customClaims, m.accessSecret, m.accessExpire)
	if err != nil {
		return nil, err
	}

	var refreshExpire time.Duration
	if inheritExpiry {
		originalExpiry := claims.RegisteredClaims.ExpiresAt.Time
		remainingTime := time.Until(originalExpiry)
		if remainingTime <= 0 {
			return nil, ErrExpiredToken
		}
		refreshExpire = remainingTime
	} else {
		refreshExpire = m.refreshExpire
	}

	refreshToken, err := m.generateToken(&customClaims, m.refreshSecret, refreshExpire)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *Manager) ValidateToken(tokenStr string, isRefresh bool) (*CustomClaims, error) {
	var secret []byte
	if isRefresh {
		secret = m.refreshSecret
	} else {
		secret = m.accessSecret
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return secret, nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return nil, ErrExpiredToken
	}

	return claims, nil
}

func (m *Manager) generateToken(claims *CustomClaims, secret []byte, duration time.Duration) (string, error) {
	now := time.Now()
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    m.issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
