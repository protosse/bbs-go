package services

import (
	"bbs-go/models"
	"bbs-go/repositories"
	"bbs-go/services/auth"

	"bbs-go/middleware/jauth"
	"github.com/kataras/iris/v12"
)

var User = newUserService()

func newUserService() *user {
	return &user{}
}

type user struct {
}

func (s *user) GetBy(id int64) *models.User {
	return repositories.User.Get(DB(), id)
}

func (s *user) GetByUsername(username string) *models.User {
	return repositories.User.GetByUsername(DB(), username)
}

func (s *user) GetByEmail(email string) *models.User {
	return repositories.User.GetByEmail(DB(), email)
}

func (s *user) Create(m *models.User) (err error) {
	return repositories.User.Create(DB(), m)
}

func (s *user) GetCurrent(ctx iris.Context) *models.User {
	userId := ctx.Values().Get("userId").(int64)
	return s.GetBy(userId)
}

func (s *user) CreateToken(userId int64) (token *jauth.Token, err error) {
	token, err = jauth.CreateToken(userId)
	if err != nil {
		return nil, err
	}

	err = auth.Driver().ToCache(userId, token)
	if err != nil {
		return nil, err
	}
	return
}
