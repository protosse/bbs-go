package api

import (
	"bbs-go/common"
	"bbs-go/middleware"
	"bbs-go/models"
	"bbs-go/services"
	"bbs-go/util"
	"bbs-go/util/date"
	"bbs-go/util/str"
	"bbs-go/util/validate"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) PostSignup() *common.JsonResult {
	body := &models.PostSignupReq{}
	var err error
	if err = c.Ctx.ReadJSON(body); err != nil {
		return common.JsonTipError(err.Error())
	}
	if msg := validate.ValidFirst(body); len(msg) != 0 {
		return common.JsonTipError(msg)
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

	err = services.User.Create(user)
	if err != nil {
		return common.JsonError(common.CreateUserError)
	}

	return common.JsonSuccess()
}

func (c *LoginController) PostLogin() *common.JsonResult {
	body := &models.PostLoginReq{}
	var err error
	if err = c.Ctx.ReadJSON(body); err != nil {
		return common.JsonTipError(err.Error())
	}
	if msg := validate.ValidFirst(body); len(msg) != 0 {
		return common.JsonTipError(msg)
	}

	var user *models.User
	if str.Len(body.Username) > 0 {
		user = services.User.GetByUsername(body.Username)
	} else if str.Len(body.Email) > 0 {
		user = services.User.GetByEmail(body.Email)
	}

	if user == nil {
		return common.JsonError(common.UserNotExistError)
	}

	if !util.ValidatePassword(user.Password, body.Password) {
		return common.JsonError(common.PasswordError)
	}

	token, err := middleware.GenerateJwtToken()
	if err != nil {
		return common.JsonError(common.GenerateTokenError)
	}

	err = services.User.Login(token, user.Id)
	if err != nil {
		return common.JsonError(common.GenerateTokenError)
	}

	res := &models.PostLoginRes{
		UserResponse: models.NewUserResponseFromModel(user),
		Token:        token,
	}
	return common.JsonData(res)
}

func (c *LoginController) GetBy(id int64) *common.JsonResult {
	user := services.User.GetBy(id)
	if user == nil {
		return common.JsonError(common.UserNotExistError)
	}

	res := models.NewUserResponseFromModel(user)
	return common.JsonData(res)
}
