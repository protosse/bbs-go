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

func ValidatePassword(raw, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(raw), []byte(input))
	return err == nil
}
