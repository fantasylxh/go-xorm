package tyt

import (
	"time"
)

type TbUser struct {
	UserId     int64     `json:"user_id" xorm:"not null pk autoincr BIGINT(20)"`
	Username   string    `json:"username" xorm:"not null comment('用户名') unique VARCHAR(50)"`
	Mobile     string    `json:"mobile" xorm:"not null comment('手机号') VARCHAR(20)"`
	Password   string    `json:"password" xorm:"comment('密码') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
}
