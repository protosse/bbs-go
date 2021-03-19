package models

type ResponseModelMap interface {
	FromModel(interface{}) *ResponseModelMap
	Dic() map[string]interface{}
}

type UserResponse struct {
	Id               int64  `json:"id"`
	UserName         string `json:"userName"`         // 用户名
	Email            string `json:"email"`            // 邮箱
	EmailVerified    bool   `json:"emailVerified"`    // 邮箱是否验证
	NickName         string `json:"nickName"`         // 昵称
	Avatar           string `json:"avatar"`           // 头像
	Description      string `json:"description"`      // 个人描述
	Status           int    `json:"status"`           // 状态
	TopicCount       int    `json:"topicCount"`       // 帖子数量
	CommentCount     int    `json:"commentCount"`     // 跟贴数量
	Roles            int    `json:"roles"`            // 用户角色
	Type             int    `json:"type"`             // 用户类型
	ForbiddenEndTime int64  `json:"forbiddenEndTime"` // 禁言结束时间
	CreateTime       int64  `json:"createTime"`       // 创建时间
}

func (u UserResponse) FromModel(i interface{}) *ResponseModelMap {
	panic("implement me")
}

func (u UserResponse) Dic() map[string]interface{} {
	panic("implement me")
}
