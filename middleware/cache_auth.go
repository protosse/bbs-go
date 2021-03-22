package middleware

import (
	"bbs-go/common"
	"bbs-go/services/auth"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func CacheAccessAuth(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	accessUuid := token["accessUuid"].(string)
	userId, err := auth.Driver().GetUserId(accessUuid)
	if err != nil {
		_, _ = ctx.JSON(common.JsonError(common.TokenError))
		return
	}
	ctx.Values().Set("userId", userId)
	ctx.Next()
}

func CacheRefreshAuth(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	refreshUuid := token["refreshUuid"].(string)
	userId, err := auth.Driver().GetUserId(refreshUuid)
	if err != nil {
		_, _ = ctx.JSON(common.JsonError(common.TokenError))
		return
	}
	ctx.Values().Set("userId", userId)
	ctx.Next()
}
