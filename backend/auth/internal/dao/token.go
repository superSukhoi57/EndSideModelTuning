package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenDAO struct {
	rdb *redis.Client
}

func NewTokenDAO(rdb *redis.Client) *TokenDAO {
	return &TokenDAO{rdb: rdb}
}

func (d *TokenDAO) SaveAccessToken(ctx context.Context, userID int64, token string, expire time.Duration) error {
	key := KeyAccessToken + fmt.Sprintf("%d", userID)
	if err := d.rdb.Set(ctx, key, token, expire).Err(); err != nil {
		return fmt.Errorf("failed to save access token: %w", err)
	}
	return nil
}

func (d *TokenDAO) SaveRefreshToken(ctx context.Context, userID int64, token string, expire time.Duration) error {
	key := KeyRefreshToken + fmt.Sprintf("%d", userID)
	if err := d.rdb.Set(ctx, key, token, expire).Err(); err != nil {
		return fmt.Errorf("failed to save refresh token: %w", err)
	}
	return nil
}

func (d *TokenDAO) GetAccessToken(ctx context.Context, userID int64) (string, error) {
	key := KeyAccessToken + fmt.Sprintf("%d", userID)
	val, err := d.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", fmt.Errorf("failed to get access token: %w", err)
	}
	return val, nil
}

func (d *TokenDAO) GetRefreshToken(ctx context.Context, userID int64) (string, error) {
	key := KeyRefreshToken + fmt.Sprintf("%d", userID)
	val, err := d.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", fmt.Errorf("failed to get refresh token: %w", err)
	}
	return val, nil
}

func (d *TokenDAO) DeleteAccessToken(ctx context.Context, userID int64) error {
	key := KeyAccessToken + fmt.Sprintf("%d", userID)
	if err := d.rdb.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete access token: %w", err)
	}
	return nil
}

func (d *TokenDAO) DeleteRefreshToken(ctx context.Context, userID int64) error {
	key := KeyRefreshToken + fmt.Sprintf("%d", userID)
	if err := d.rdb.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}
	return nil
}

func (d *TokenDAO) DeleteAllTokens(ctx context.Context, userID int64) error {
	if err := d.DeleteAccessToken(ctx, userID); err != nil {
		return err
	}
	if err := d.DeleteRefreshToken(ctx, userID); err != nil {
		return err
	}
	return nil
}
