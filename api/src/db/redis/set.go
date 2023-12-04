package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
)

func (r *query) SaveToken(ctx context.Context, model tables.Access) error {
	ttl, err := strconv.Atoi(r.config.AccessTokenExpiredMin)
	if err != nil {
		r.log.Error("SaveToken -> AccessTokenExpiredMin", "error", err)
		return err
	}
	err = r.set(ctx, model.Token, model.Token, time.Minute*time.Duration(ttl))
	if err != nil {
		r.log.Error("SaveToken -> Token", "error", err)
		return err
	}

	ttl, err = strconv.Atoi(r.config.RefreshTokenExpiredMin)
	if err != nil {
		r.log.Error("SaveToken -> RefreshTokenExpiredMin", "error", err)
		return err
	}
	err = r.set(ctx, model.TokenRefresh, model.TokenRefresh, time.Minute*time.Duration(ttl))
	if err != nil {
		r.log.Error("SaveToken -> TokenRefresh", "error", err)
		return err
	}
	return nil
}

func (r *query) GetToken(ctx context.Context, token string) (bool, error) {
	_, err := r.get(ctx, token)
	if err != nil {
		r.log.Error("GetToken ", "error", err)
		return false, err
	}

	return true, nil
}

func (r *query) RemoveToken(ctx context.Context, token string) error {
	err := r.del(ctx, token)
	if err != nil {
		r.log.Error("RemoveToken ", "error", err)
		return err
	}

	return nil
}

func (r *query) set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *query) get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *query) del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
