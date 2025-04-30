package usecase

import "context"

type RedisUsecase interface {
	Save(ctx context.Context, userID string, jti string, ttlSec int) error
	Match(ctx context.Context, userID string, jti string) (bool, error)
	Delete(ctx context.Context, userID string) error
}
