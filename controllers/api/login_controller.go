package api

import (
	"bbs-go/common"
	"bbs-go/middleware"
	"bbs-go/middleware/jauth"
	"bbs-go/models"
	"bbs-go/services"
	"bbs-go/util"
	"bbs-go/util/date"
	"bbs-go/util/logging"
	"bbs-go/util/str"
	"bbs-go/util/validate"
	"github.com/kataras/iris/v12/mvc"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("Post", "/refreshToken", "RefreshToken", jauth.Refresh(), middleware.CacheRefreshAuth)
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

	token, err := services.User.CreateToken(user.Id)
	if err != nil {
		logging.Errorf("%v", err)
		return common.JsonError(common.CreateTokenError)
	}

	res := &models.PostLoginRes{
		UserResponse: models.NewUserResponseFromModel(user),
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
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

func (c *LoginController) RefreshToken() *common.JsonResult {
	userId := c.Ctx.Values().Get("userId").(int64)
	token, err := services.User.CreateToken(userId)
	if err != nil {
		return common.JsonError(common.CreateTokenError)
	}
	res := map[string]string{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
	}
	return common.JsonData(res)
}
