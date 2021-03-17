package param

type PostSignupReq struct {
	Username UserName `json:"username"`
	Email    Email    `json:"email"`
	Password Password `json:"password"`
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
