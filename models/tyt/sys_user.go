package tyt

import (
	"time"
)

type SysUser struct {
	UserId       int64     `json:"user_id" xorm:"not null pk autoincr BIGINT(20)"`
	Username     string    `json:"username" xorm:"not null comment('用户名') unique VARCHAR(50)"`
	Password     string    `json:"password" xorm:"comment('密码') VARCHAR(100)"`
	Email        string    `json:"email" xorm:"comment('邮箱') VARCHAR(100)"`
	Mobile       string    `json:"mobile" xorm:"comment('手机号') VARCHAR(100)"`
	Status       int       `json:"status" xorm:"comment('状态  0：禁用   1：正常') TINYINT(4)"`
	CreateUserId int64     `json:"create_user_id" xorm:"comment('创建者ID') BIGINT(20)"`
	CreateTime   time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
}
