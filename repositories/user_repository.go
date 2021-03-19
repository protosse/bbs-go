package repositories

import (
	"bbs-go/models"
	"gorm.io/gorm"
)

var User = newUserRepository()

type user struct {
}

func newUserRepository() *user {
	return &user{}
}

func (r *user) Get(db *gorm.DB, id int64) *models.User {
	ret := &models.User{}
	if err := db.First(ret, id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *user) Take(db *gorm.DB, where ...interface{}) *models.User {
	ret := &models.User{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *user) Create(db *gorm.DB, m *models.User) error {
	err := db.Create(m).Error
	return err
}

func (r *user) GetByUsername(db *gorm.DB, username string) *models.User {
	return r.Take(db, "user_name = ?", username)
}

func (r *user) GetByEmail(db *gorm.DB, email string) *models.User {
	return r.Take(db, "email = ?", email)
}
