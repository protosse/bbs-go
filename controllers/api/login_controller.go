package api

import (
	"bbs-go/common"
	"bbs-go/controllers/param"
	"bbs-go/models"
	"bbs-go/services"
	"bbs-go/util"
	"bbs-go/util/date"
	"bbs-go/util/str"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) PostSignup() *common.JsonResult {
	body := &param.PostSignupReq{}
	if err := c.Ctx.ReadJSON(body); err != nil {
		return common.JsonTipError(str.RemoveUnmarshalerDecoder(err.Error()))
	}

	if services.User.GetByUsername(body.Username.Value) != nil {
		return common.JsonError(common.UsernameExistError)
	}

	if services.User.GetByEmail(body.Email.Value) != nil {
		return common.JsonError(common.EmailExistError)
	}

	user := &models.User{
		UserName:   services.SqlNullString(body.Username.Value),
		Email:      services.SqlNullString(body.Email.Value),
		Password:   util.EncodePassword(body.Password.Value),
		CreateTime: date.NowTimestamp(),
		UpdateTime: date.NowTimestamp(),
	}

	err := services.User.Create(user)
	if err != nil {
		return common.JsonError(common.CreateUserError)
	}

	return common.JsonSuccess()
}
