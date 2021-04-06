package auth

import (
	"bbs-go/middleware/jauth"
	"time"
)

type RedisAuth struct {
	Conn *RedisCluster
}

func (r RedisAuth) ToCache(userId int64, token *jauth.Token) error {
	now := time.Now()
	_, err := r.Conn.Set(token.AccessUuid, userId, time.Unix(token.AccessTokenExpires, 0).Sub(now))
	if err != nil {
		return err
	}
	_, err = r.Conn.Set(token.RefreshUuid, userId, time.Unix(token.RefreshTokenExpires, 0).Sub(now))
	if err != nil {
		return err
	}
	return nil
}

func (r RedisAuth) GetUserId(uuid string) (int64, error) {
	userId, err := r.Conn.GetKey(uuid)
	if err != nil {
		return 0, err
	}
	return userId.(int64), nil
}

func (r RedisAuth) DelCache(uuid string) (int64, error) {
	userId, err := r.GetUserId(uuid)
	if err != nil {
		return 0, err
	}
	_, err = r.Conn.Del(uuid)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (r RedisAuth) Close() {
	r.Conn.Close()
}

func NewRedisAuth() *RedisAuth {
	return &RedisAuth{
		Conn: GetRedisClusterClient(),
	}
}
