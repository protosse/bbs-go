package api

import (
	"github.com/kataras/iris/v12"
)

type TokenController struct {
	Ctx *iris.Context
}

//func (c *TokenController)Get() *common.JsonResult {
//	token, err := middleware.GenerateJwtToken()
//	if err != nil {
//		return common.JsonError(common.GenerateTokenError)
//	}
//
//	res := &models.PostLoginRes{
//		UserResponse: models.NewUserResponseFromModel(user),
//		Token: token,
//	}
//	return common.JsonData(res)
//}
