package auth

import (
	"bbs-go/common/constants"
	"time"

	"github.com/patrickmn/go-cache"
)

var localCache *cache.Cache

type LocalAuth struct {
	Cache *cache.Cache
}

func cureKey(key string) string {
	return SessionTokenPrefix + key
}

func (l *LocalAuth) ToCache(token string, id int64) error {
	key := cureKey(token)
	session := &Session{
		UserId:    id,
		CreatDate: time.Now().Unix(),
	}
	l.Cache.Set(key, session, 0)
	return nil
}

func (l *LocalAuth) GetSession(token string) (*Session, error) {
	get, found := l.Cache.Get(cureKey(token))
	if !found {
		return nil, ErrTokenInvalid
	}
	return get.(*Session), nil
}

func (l *LocalAuth) Close() {}

func NewLocalAuth() *LocalAuth {
	if localCache == nil {
		localCache = cache.New(constants.DefaultTokenExpireHour*time.Hour, 20*time.Minute)
	}
	return &LocalAuth{
		Cache: localCache,
	}
}
