package api

import (
	"bbs-go/common"
	"bbs-go/models"
	"bbs-go/services"
	"bbs-go/util"
	"bbs-go/util/date"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
}

type signupPostBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *LoginController) PostSignup() *common.JsonResult {
	body := &signupPostBody{}
	if err := c.Ctx.ReadJSON(body); err != nil {
		return common.JsonTipError(err.Error())
	}

	// if !captcha.VerifyString(captchaId, captchaCode) {
	// 	return common.JsonError(common.CaptchaError)
	// }

	if err := util.IsEmail(body.Email); err != nil {
		return common.JsonTipError(err.Error())
	}

	if err := util.IsUsername(body.Username); err != nil {
		return common.JsonTipError(err.Error())
	}

	if err := util.IsPassword(body.Password); err != nil {
		return common.JsonTipError(err.Error())
	}

	if services.User.GetByUsername(body.Username) != nil {
		return common.JsonError(common.UsernameExistError)
	}

	if services.User.GetByEmail(body.Email) != nil {
		return common.JsonError(common.EmailExistError)
	}

	user := &models.User{
		UserName:   services.SqlNullString(body.Username),
		Email:      services.SqlNullString(body.Email),
		Password:   util.EncodePassword(body.Password),
		CreateTime: date.NowTimestamp(),
		UpdateTime: date.NowTimestamp(),
	}

	err := services.User.Create(user)
	if err != nil {
		return common.JsonError(common.CreateUserError)
	}

	return common.JsonData(user)
}
