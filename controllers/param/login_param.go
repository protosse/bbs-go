package param

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

type PostSignupRes struct {
	Id               int64  `json:"id"`
	UserName         string `json:"username"`
	Email            string `json:"email"`
	EmailVerified    bool   `json:"emailVerified"`
	NickName         string `json:"nickName"`
	Avatar           string `json:"avatar"`
	Description      string `json:"description"`
	Status           int    `json:"status"`
	TopicCount       int    `json:"topicCount"`
	CommentCount     int    `json:"commentCount"`
	Roles            int    `json:"roles"`
	Type             int    `json:"type"`
	ForbiddenEndTime int64  `json:"forbiddenEndTime"`
	CreateTime       int64  `json:"createTime"`
}
