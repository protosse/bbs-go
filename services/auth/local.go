package auth

import (
	"bbs-go/common/constants"
	"bbs-go/middleware/jauth"
	"time"

	"github.com/patrickmn/go-cache"
)

var localCache *cache.Cache

type LocalAuth struct {
	Cache *cache.Cache
}

func (l *LocalAuth) ToCache(userId int64, token *jauth.Token) error {
	now := time.Now()
	l.Cache.Set(token.AccessUuid, userId, time.Unix(token.AccessTokenExpires, 0).Sub(now))
	l.Cache.Set(token.RefreshUuid, userId, time.Unix(token.RefreshTokenExpires, 0).Sub(now))
	return nil
}

func (l *LocalAuth) GetUserId(uuid string) (int64, error) {
	userId, found := l.Cache.Get(uuid)
	if !found {
		return 0, ErrTokenInvalid
	}
	return userId.(int64), nil
}

func (l *LocalAuth) DelCache(uuid string) (int64, error) {
	userId, err := l.GetUserId(uuid)
	if err != nil {
		return 0, err
	}
	l.Cache.Delete(uuid)
	return userId, nil
}

func (l *LocalAuth) Close() {}

func NewLocalAuth() *LocalAuth {
	if localCache == nil {
		localCache = cache.New(constants.AccessTokenExpireHour*time.Hour, 20*time.Minute)
	}
	return &LocalAuth{
		Cache: localCache,
	}
}
