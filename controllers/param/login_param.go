package param

import "bbs-go/models"

type PostSignupReq struct {
	Username string `json:"username" validate:"required,username"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
}

type PostLoginReq struct {
	Username string `json:"username" validate:"username"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
}

type PostLoginRes struct {
	models.UserResponse
	Token string `json:"token"`
}
