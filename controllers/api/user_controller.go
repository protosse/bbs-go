package api

import (
	"bbs-go/common"
	"bbs-go/middleware"
	"bbs-go/services"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

func (c *UserController) GetCurrent() *common.JsonResult {
	user := services.User.GetCurrent(c.Ctx)
	if user == nil {
		return common.JsonError(common.UserNotExistError)
	}
	return common.JsonData(user)
}

func (c *UserController) GetToken() *common.JsonResult {
	token, err := middleware.GenerateJwtToken()
	if err != nil {
		return common.JsonError(common.GenerateTokenError)
	}
	return common.JsonData(map[string]string{
		"token": token,
	})
}
