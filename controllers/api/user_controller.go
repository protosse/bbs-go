package api

import (
	"bbs-go/common"
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
