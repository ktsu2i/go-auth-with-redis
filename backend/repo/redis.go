package repo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	rc *redis.Client
}

func NewRedisRepo(rc *redis.Client) *RedisRepo {
	return &RedisRepo{rc: rc}
}

func (r *RedisRepo) Save(ctx context.Context, userID string, jti string, ttlSec int) error {
	return r.rc.Set(ctx, userID, jti, time.Duration(ttlSec)*time.Second).Err()
}

func (r *RedisRepo) Match(ctx context.Context, userID string, jti string) (bool, error) {
	val, err := r.rc.Get(ctx, userID).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val == jti, nil
}

func (r *RedisRepo) Delete(ctx context.Context, userID string) error {
	return r.rc.Del(ctx, userID).Err()
}
