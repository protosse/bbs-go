package services

import (
	"bbs-go/common/constants"
	"bbs-go/models"
	"bbs-go/util"
	"bbs-go/util/date"
	"time"
)

var User = newUserService()

func newUserService() *user {
	return &user{}
}

type user struct {
}

func (s *user) GetByUsername(username string) *models.User {
	ret := &models.User{}
	if err := DB().Take(ret, "username = ?", username); err != nil {
		return nil
	}
	return ret
}

func (s *user) GetByEmail(email string) *models.User {
	ret := &models.User{}
	if err := DB().Take(ret, "email = ?", email); err != nil {
		return nil
	}
	return ret
}

func (s *user) Create(u *models.User) (err error) {
	err = DB().Create(u).Error
	return
}

func (s *user) CreateToken(userId int64) (string, error) {
	token := util.UUID()
	expiredAt := time.Now().Add(time.Hour * time.Duration(24) * constants.DefaultTokenExporeDays)
	userToken := &models.UserToken{
		Token:      token,
		UserId:     userId,
		ExpiredAt:  date.Timestamp(expiredAt),
		Status:     constants.StatusOk,
		CreateTime: date.NowTimestamp(),
	}
	err := DB().Create(userToken).Error
	if err != nil {
		return "", err
	}
	return token, nil
}