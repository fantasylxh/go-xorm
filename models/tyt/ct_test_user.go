package tyt

import (
	"time"
)

type CtTestUser struct {
	Id         string    `json:"id" xorm:"not null pk default '' comment('id') VARCHAR(64)"`
	CreateDate time.Time `json:"create_date" xorm:"comment('创建时间') DATETIME"`
	ModifyDate time.Time `json:"modify_date" xorm:"comment('修改时间') DATETIME"`
	Name       string    `json:"name" xorm:"comment('姓名') VARCHAR(64)"`
	Sex        string    `json:"sex" xorm:"comment('性别') VARCHAR(255)"`
	Tata       string    `json:"tata" xorm:"comment('测试') VARCHAR(255)"`
	IsUpdate   int       `json:"is_update" xorm:"comment('更新与否') TINYINT(1)"`
}
