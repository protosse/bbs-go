package auth

import (
	"bbs-go/common/config"
	"bbs-go/middleware/jauth"
	"errors"
)

var (
	ErrTokenInvalid = errors.New("无效的token")
)

var authDriver Authentication

func Driver() Authentication {
	if authDriver != nil {
		return authDriver
	}

	switch config.Global.Cache.Driver {
	default:
		return NewLocalAuth()
	}
}

type Authentication interface {
	ToCache(userId int64, token *jauth.Token) error
	GetUserId(uuid string) (int64, error)
	DelCache(uuid string) (int64, error)
	Close()
}
