package models

import "database/sql"

var Models = []interface{}{
	&User{}, &UserToken{}, &Topic{}, &TopicNode{}, &Comment{},
}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement"`
}

type User struct {
	Model
	UserName         sql.NullString `gorm:"size:32;unique"`                 // 用户名
	Email            sql.NullString `gorm:"size:128;unique"`                // 邮箱
	EmailVerified    bool           `gorm:"not null;default:false"`         // 邮箱是否验证
	NickName         string         `gorm:"size:16"`                        // 昵称
	Avatar           string         `gorm:"type:text"`                      // 头像
	Password         string         `gorm:"type:text"`                      // 密码
	Description      string         `gorm:"type:text"`                      // 个人描述
	Status           int            `gorm:"index:idx_user_status;not null"` // 状态
	TopicCount       int            `gorm:"not null"`                       // 帖子数量
	CommentCount     int            `gorm:"not null"`                       // 跟贴数量
	Roles            int            `gorm:"not null"`                       // 用户角色
	Type             int            `gorm:"not null"`                       // 用户类型
	ForbiddenEndTime int64          `gorm:"not null;default:0"`             // 禁言结束时间
	CreateTime       int64          // 创建时间
	UpdateTime       int64          // 更新时间
}

type UserToken struct {
	Model
	Token      string `gorm:"size:32;unique;not null"`
	UserId     int64  `gorm:"not null;index:idx_user_token_user_id;"`
	ExpiredAt  int64  `gorm:"not null"`
	Status     int    `gorm:"not null;index:idx_user_token_status"`
	CreateTime int64  `gorm:"not null"`
}
type Topic struct {
	Model
	NodeId            int64  `gorm:"not null;index:idx_node_id;"`       // 节点编号
	UserId            int64  `gorm:"not null;index:idx_topic_user_id;"` // 用户
	Title             string `gorm:"size:128"`                          // 标题
	Content           string `gorm:"type:longtext"`                     // 内容
	ImageList         string `gorm:"type:longtext"`                     // 图片
	ViewCount         int64  `gorm:"not null"`                          // 查看数量
	CommentCount      int64  `gorm:"not null"`                          // 跟帖数量
	Status            int    `gorm:"index:idx_topic_status;"`           // 状态：0：正常、1：删除
	LastCommentTime   int64  `gorm:"index:idx_topic_last_comment_time"` // 最后回复时间
	LastCommentUserId int64  // 最后回复用户
	CreateTime        int64  `gorm:"index:idx_topic_create_time"` // 创建时间
	ExtraData         string `gorm:"type:text"`                   // 扩展数据
}
type TopicNode struct {
	Model
	Name        string `gorm:"size:32;unique"` // 名称
	Description string // 描述
	SortNo      int    `gorm:"index:idx_sort_no"` // 排序编号
	Status      int    `gorm:"not null"`          // 状态
	CreateTime  int64  // 创建时间
}
type Comment struct {
	Model
	UserId      int64  `gorm:"index:idx_comment_user_id;not null"`     // 用户编号
	EntityType  string `gorm:"index:idx_comment_entity_type;not null"` // 被评论实体类型
	EntityId    int64  `gorm:"index:idx_comment_entity_id;not null"`   // 被评论实体编号
	Content     string `gorm:"type:text;not null"`                     // 内容
	ContentType string `gorm:"type:varchar(32);not null"`              // 内容类型：markdown、html
	QuoteId     int64  `gorm:"not null" `                              // 引用的评论编号
	Status      int    `gorm:"int;index:idx_comment_status"`           // 状态：0：待审核、1：审核通过、2：审核失败、3：已发布
	CreateTime  int64  // 创建时间
}
