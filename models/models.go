package models

import "database/sql"

var Models = []interface{}{
	&User{}, &UserToken{}, &Topic{}, &TopicNode{}, &Comment{},
}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

type User struct {
	Model
	UserName         sql.NullString `gorm:"size:32;unique" json:"username" form:"username"`                     // 用户名
	Email            sql.NullString `gorm:"size:128;unique" json:"email" form:"email"`                          // 邮箱
	EmailVerified    bool           `gorm:"not null;default:false" json:"emailVerified" form:"emailVerified"`   // 邮箱是否验证
	NickName         string         `gorm:"size:16" json:"nickname" form:"nickname"`                            // 昵称
	Avatar           string         `gorm:"type:text" json:"avatar" form:"avatar"`                              // 头像
	Password         string         `gorm:"type:text" json:"password" form:"password"`                          // 密码
	Description      string         `gorm:"type:text" json:"description" form:"description"`                    // 个人描述
	Status           int            `gorm:"index:idx_user_status;not null" json:"status" form:"status"`         // 状态
	TopicCount       int            `gorm:"not null" json:"topicCount" form:"topicCount"`                       // 帖子数量
	CommentCount     int            `gorm:"not null" json:"commentCount" form:"commentCount"`                   // 跟贴数量
	Roles            int            `gorm:"not null" json:"roles" form:"roles"`                                 // 用户角色
	Type             int            `gorm:"not null" json:"type" form:"type"`                                   // 用户类型
	ForbiddenEndTime int64          `gorm:"not null;default:0" json:"forbiddenEndTime" form:"forbiddenEndTime"` // 禁言结束时间
	CreateTime       int64          `json:"createTime" form:"createTime"`                                       // 创建时间
	UpdateTime       int64          `json:"updateTime" form:"updateTime"`                                       // 更新时间
}

type UserToken struct {
	Model
	Token      string `gorm:"size:32;unique;not null" json:"token" form:"token"`
	UserId     int64  `gorm:"not null;index:idx_user_token_user_id;" json:"userId" form:"userId"`
	ExpiredAt  int64  `gorm:"not null" json:"expiredAt" form:"expiredAt"`
	Status     int    `gorm:"not null;index:idx_user_token_status" json:"status" form:"status"`
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`
}
type Topic struct {
	Model
	NodeId            int64  `gorm:"not null;index:idx_node_id;" json:"nodeId" form:"nodeId"`                         // 节点编号
	UserId            int64  `gorm:"not null;index:idx_topic_user_id;" json:"userId" form:"userId"`                   // 用户
	Title             string `gorm:"size:128" json:"title" form:"title"`                                              // 标题
	Content           string `gorm:"type:longtext" json:"content" form:"content"`                                     // 内容
	ImageList         string `gorm:"type:longtext" json:"imageList" form:"imageList"`                                 // 图片
	ViewCount         int64  `gorm:"not null" json:"viewCount" form:"viewCount"`                                      // 查看数量
	CommentCount      int64  `gorm:"not null" json:"commentCount" form:"commentCount"`                                // 跟帖数量
	Status            int    `gorm:"index:idx_topic_status;" json:"status" form:"status"`                             // 状态：0：正常、1：删除
	LastCommentTime   int64  `gorm:"index:idx_topic_last_comment_time" json:"lastCommentTime" form:"lastCommentTime"` // 最后回复时间
	LastCommentUserId int64  `json:"lastCommentUserId" form:"lastCommentUserId"`                                      // 最后回复用户
	CreateTime        int64  `gorm:"index:idx_topic_create_time" json:"createTime" form:"createTime"`                 // 创建时间
	ExtraData         string `gorm:"type:text" json:"extraData" form:"extraData"`                                     // 扩展数据
}
type TopicNode struct {
	Model
	Name        string `gorm:"size:32;unique" json:"name" form:"name"`        // 名称
	Description string `json:"description" form:"description"`                // 描述
	SortNo      int    `gorm:"index:idx_sort_no" json:"sortNo" form:"sortNo"` // 排序编号
	Status      int    `gorm:"not null" json:"status" form:"status"`          // 状态
	CreateTime  int64  `json:"createTime" form:"createTime"`                  // 创建时间
}
type Comment struct {
	Model
	UserId      int64  `gorm:"index:idx_comment_user_id;not null" json:"userId" form:"userId"`             // 用户编号
	EntityType  string `gorm:"index:idx_comment_entity_type;not null" json:"entityType" form:"entityType"` // 被评论实体类型
	EntityId    int64  `gorm:"index:idx_comment_entity_id;not null" json:"entityId" form:"entityId"`       // 被评论实体编号
	Content     string `gorm:"type:text;not null" json:"content" form:"content"`                           // 内容
	ContentType string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`            // 内容类型：markdown、html
	QuoteId     int64  `gorm:"not null"  json:"quoteId" form:"quoteId"`                                    // 引用的评论编号
	Status      int    `gorm:"int;index:idx_comment_status" json:"status" form:"status"`                   // 状态：0：待审核、1：审核通过、2：审核失败、3：已发布
	CreateTime  int64  `json:"createTime" form:"createTime"`                                               // 创建时间
}
