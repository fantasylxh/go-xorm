package tyt

import (
	"time"
)

type UserToken struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username   string    `json:"username" xorm:"index(index_name) VARCHAR(64)"`
	Token      string    `json:"token" xorm:"index(index_name) VARCHAR(64)"`
	CreateDate time.Time `json:"create_date" xorm:"DATETIME"`
	ModifyDate time.Time `json:"modify_date" xorm:"DATETIME"`
}
