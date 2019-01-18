package tyt

type SysUserRole struct {
	Id     int64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId int64 `json:"user_id" xorm:"comment('用户ID') BIGINT(20)"`
	RoleId int64 `json:"role_id" xorm:"comment('角色ID') BIGINT(20)"`
}
