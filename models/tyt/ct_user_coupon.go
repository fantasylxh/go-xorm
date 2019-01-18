package tyt

import (
	"time"
)

type CtUserCoupon struct {
	Id         string    `json:"id" xorm:"not null pk VARCHAR(255)"`
	CreateDate time.Time `json:"create_date" xorm:"DATETIME"`
	ModifyDate time.Time `json:"modify_date" xorm:"DATETIME"`
	UserId     string    `json:"user_id" xorm:"VARCHAR(255)"`
	CouponId   int       `json:"coupon_id" xorm:"INT(11)"`
	Status     string    `json:"status" xorm:"VARCHAR(255)"`
	LimitDate  time.Time `json:"limit_date" xorm:"DATETIME"`
	Price      string    `json:"price" xorm:"DECIMAL(11,2)"`
	UseDate    time.Time `json:"use_date" xorm:"DATETIME"`
	OrdersId   string    `json:"orders_id" xorm:"VARCHAR(255)"`
	Minimum    string    `json:"minimum" xorm:"DECIMAL(11,2)"`
}
