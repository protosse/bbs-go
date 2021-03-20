package auth

import (
	"bbs-go/common/config"
	"errors"
)

var (
	SessionTokenPrefix = "GB:"
	ErrTokenInvalid    = errors.New("无效的token")
)

var authDriver Authentication

func Driver() Authentication {
	if authDriver != nil {
		return authDriver
	}

	switch config.Config.Cache.Driver {
	default:
		return NewLocalAuth()
	}
}

type Authentication interface {
	ToCache(token string, id int64) error
	GetSession(token string) (*Session, error)
	Close()
}

type Session struct {
	UserId    int64  `json:"userId" redis:"userId"`
	CreatDate int64  `json:"creatDate" redis:"creatDate"`
	ExpiresIn int64  `json:"expiresIn" redis:"expiresIn"`
	Scope     uint64 `json:"scope" redis:"scope"`
}
