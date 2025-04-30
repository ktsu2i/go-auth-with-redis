package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	db *redis.Client
}

func NewRedisRepo(db *redis.Client) *RedisRepo {
	return &RedisRepo{db: db}
}

func key(userID string) string {
	return fmt.Sprintf("rt:%s", userID)
}

func (r *RedisRepo) Save(ctx context.Context, userID string, jti string, ttlSec int) error {
	return r.db.Set(ctx, key(userID), jti, time.Duration(ttlSec)*time.Second).Err()
}

func (r *RedisRepo) Match(ctx context.Context, userID string, jti string) (bool, error) {
	val, err := r.db.Get(ctx, key(userID)).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val == jti, nil
}

func (r *RedisRepo) Delete(ctx context.Context, userID string) error {
	return r.db.Del(ctx, key(userID)).Err()
}
