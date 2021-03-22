package jauth

import (
	"bbs-go/common"
	"bbs-go/common/constants"
	"bbs-go/util/str"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"time"
)

var (
	accessSignKey  = []byte("HS2JDFKhu7Y1av7b")
	refreshSignKey = []byte("HS2JDFKhu7Y1av2b")
	signMethod     = jwt.SigningMethodHS256
)

type Token struct {
	AccessToken         string
	RefreshToken        string
	AccessUuid          string
	RefreshUuid         string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

func CreateToken(userId int64) (*Token, error) {
	var err error
	t := &Token{}

	t.AccessUuid = str.UUID()
	t.AccessTokenExpires = time.Now().Add(time.Hour * constants.AccessTokenExpireHour).Unix()
	t.AccessToken, err = jwt.NewTokenWithClaims(signMethod, jwt.MapClaims{
		"authorized": true,
		"userId":     userId,
		"accessUuid": t.AccessUuid,
		"exp":        t.AccessTokenExpires,
	}).SignedString(accessSignKey)

	if err != nil {
		return nil, err
	}

	t.RefreshUuid = str.UUID()
	t.RefreshTokenExpires = time.Now().Add(time.Hour * constants.RefreshTokenExpireHour).Unix()
	t.RefreshToken, err = jwt.NewTokenWithClaims(signMethod, jwt.MapClaims{
		"authorized": true,
		"userId":     userId,
		"accessUuid": t.RefreshUuid,
		"exp":        t.RefreshTokenExpires,
	}).SignedString(refreshSignKey)

	if err != nil {
		return nil, err
	}

	return t, err
}

func Access() context.Handler {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return accessSignKey, nil
		},
		SigningMethod: signMethod,
		ErrorHandler: func(context iris.Context, err error) {
			_, _ = context.JSON(common.JsonError(common.TokenError))
		},
	}).Serve
}

func Refresh() context.Handler {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return refreshSignKey, nil
		},
		SigningMethod: signMethod,
		ErrorHandler: func(context iris.Context, err error) {
			_, _ = context.JSON(common.JsonError(common.TokenError))
		},
	}).Serve
}
