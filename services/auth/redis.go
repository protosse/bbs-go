package auth

import (
	"bbs-go/common/config"
	"bbs-go/middleware/jauth"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisAuth struct {
	Client *redis.Client
	ctx    context.Context
}

func (r RedisAuth) ToCache(userId int64, token *jauth.Token) error {
	now := time.Now()
	var err error
	err = r.Client.Set(r.ctx, token.AccessUuid, userId, time.Unix(token.AccessTokenExpires, 0).Sub(now)).Err()
	if err != nil {
		return err
	}
	err = r.Client.Set(r.ctx, token.RefreshUuid, userId, time.Unix(token.RefreshTokenExpires, 0).Sub(now)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisAuth) GetUserId(uuid string) (int64, error) {
	userId, err := r.Client.Get(r.ctx, uuid).Int64()
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (r RedisAuth) DelCache(uuid string) (int64, error) {
	userId, err := r.GetUserId(uuid)
	if err != nil {
		return 0, err
	}
	err = r.Client.Del(r.ctx, uuid).Err()
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (r RedisAuth) Close() {
	_ = r.Client.Close()
}

func NewRedisAuth(config *config.Config) *RedisAuth {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})
	return &RedisAuth{
		ctx:    context.Background(),
		Client: client,
	}
}
