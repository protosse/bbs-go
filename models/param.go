package models

type PostSignupReq struct {
	Username string `json:"username" validate:"required,userNameIgnoreEmpty"`
	Email    string `json:"email" validate:"required,emailIgnoreEmpty"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
}

type PostLoginReq struct {
	Username string `json:"username" validate:"userNameIgnoreEmpty"`
	Email    string `json:"email" validate:"emailIgnoreEmpty"`
	Password string `json:"password" validate:"required,gte=6,lte=30"`
}

type PostLoginRes struct {
	*UserResponse
	Token string `json:"token"`
}
