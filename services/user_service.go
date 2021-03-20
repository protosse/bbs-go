package services

import (
	"bbs-go/models"
	"bbs-go/repositories"
	"bbs-go/services/auth"

	"github.com/iris-contrib/middleware/jwt"
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
	token := ctx.Values().Get("jwt").(*jwt.Token).Raw
	session, err := auth.Driver().GetSession(token)
	if err != nil {
		return nil
	}

	return s.GetBy(session.UserId)
}

func (s *user) Login(token string, id int64) (err error) {
	return auth.Driver().ToCache(token, id)
}
