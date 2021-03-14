package util

import "golang.org/x/crypto/bcrypt"

func EncodePassword(raw string) (hash string) {
	bt, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		hash = raw
	} else {
		hash = string(bt)
	}
	return
}
