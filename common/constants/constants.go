package constants

const (
	AccessTokenExpireHour  = 1      //用户登录token默认有效期(hour)
	RefreshTokenExpireHour = 24 * 7 //用户登录token默认有效期(hour)
)

// 用户角色
const (
	RoleUser  = iota // 用户
	RoleOwner        //站长
	RoleAdmin        //管理员
)

// 状态
const (
	StatusOk      = 0 // 正常
	StatusDeleted = 1 // 删除
	StatusPending = 2 // 待审核
)
