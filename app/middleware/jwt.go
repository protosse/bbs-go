package middleware

import (
	"bbs-go/common"
	"bbs-go/common/constants"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"time"
)

var (
	signKey    = []byte(constants.SignKey)
	signMethod = jwt.SigningMethodHS256
)

func GenerateJwtToken() (string, error) {
	token := jwt.NewTokenWithClaims(signMethod, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * constants.DefaultTokenExpireHour).Unix(),
	})
	str, err := token.SignedString(signKey)
	return str, err
}

func JwtHandler() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signKey, nil
		},
		SigningMethod: signMethod,
		ErrorHandler: func(context iris.Context, err error) {
			_, _ = context.JSON(common.JsonError(common.TokenError))
		},
	})
}
