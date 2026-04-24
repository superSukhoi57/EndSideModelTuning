package dao

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type StateDAO struct {
	rdb *redis.Client
}

func NewStateDAO(rdb *redis.Client) *StateDAO {
	return &StateDAO{
		rdb: rdb,
	}
}

func (d *StateDAO) GenerateAndStoreState(ctx context.Context, ttl time.Duration) (string, error) {

	state, err := generateState()
	logx.Infof("生成state %s 到后端redis  ttl: %d", state, ttl)
	if err != nil {
		return "", fmt.Errorf("failed to generate state: %w", err)
	}

	key := fmt.Sprintf("auth:state:%s", state)
	if err := d.rdb.Set(ctx, key, "1", ttl).Err(); err != nil {
		return "", fmt.Errorf("failed to store state in redis: %w", err)
	}

	return state, nil
}

func (d *StateDAO) ValidateState(ctx context.Context, state string) (bool, error) {
	key := fmt.Sprintf("auth:state:%s", state)
	exists, err := d.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to validate state: %w", err)
	}

	if exists > 0 {
		d.rdb.Del(ctx, key)
		return true, nil
	}

	return false, nil
}

func generateState() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
