package tyt

import (
	"gameCityService/global"
	"time"
)

type User struct {
	Id                  string    `json:"id" xorm:"not null pk VARCHAR(255)"`
	Username            string    `json:"username" xorm:"unique VARCHAR(100)"`
	BuyDate             time.Time `json:"buy_date" xorm:"DATETIME"`
	CreateDate          time.Time `json:"create_date" xorm:"DATETIME"`
	TimeLeft            int       `json:"time_left" xorm:"default 0 INT(11)"`
	Phone               string    `json:"phone" xorm:"VARCHAR(11)"`
	Status              string    `json:"status" xorm:"CHAR(1)"`
	IsBlack             string    `json:"is_black" xorm:"CHAR(1)"`
	UserSource          string    `json:"user_source" xorm:"CHAR(1)"`
	Password            string    `json:"password" xorm:"not null VARCHAR(300)"`
	Level               string    `json:"level" xorm:"CHAR(1)"`
	Token               string    `json:"token" xorm:"VARCHAR(50)"`
	Fmcode              string    `json:"fmcode" xorm:"VARCHAR(10)"`
	LoginDate           time.Time `json:"login_date" xorm:"DATETIME"`
	Ip                  string    `json:"ip" xorm:"VARCHAR(255)"`
	Version             string    `json:"version" xorm:"VARCHAR(255)"`
	VipEndTime          time.Time `json:"vip_end_time" xorm:"DATETIME"`
	ShareCode           string    `json:"share_code" xorm:"VARCHAR(255)"`
	LeftCount           int       `json:"left_count" xorm:"default 0 INT(11)"`
	IsRegisterSend      int       `json:"is_register_send" xorm:"INT(1)"`
	IsVipSend           int       `json:"is_vip_send" xorm:"INT(1)"`
	FreeStartTime       time.Time `json:"free_start_time" xorm:"DATETIME"`
	SignTime            time.Time `json:"sign_time" xorm:"DATETIME"`
	IsEmailSend         int       `json:"is_email_send" xorm:"INT(1)"`
	CommissionAmount    string    `json:"commission_amount" xorm:"DECIMAL(11,2)"`
	ModifyDate          time.Time `json:"modify_date" xorm:"DATETIME"`
	SendDate            time.Time `json:"send_date" xorm:"DATETIME"`
	LimitSendTime       int       `json:"limit_send_time" xorm:"default 0 INT(6)"`
	PhoneEndTime        time.Time `json:"phone_end_time" xorm:"DATETIME"`
	ShadowsocksPassword string    `json:"shadowsocks_password" xorm:"VARCHAR(255)"`
	ShadowsocksPort     string    `json:"shadowsocks_port" xorm:"VARCHAR(255)"`
	Depname             string    `json:"depname" xorm:"comment('机构') VARCHAR(50)"`
	Note                string    `json:"note" xorm:"VARCHAR(255)"`
	Type                string    `json:"type" xorm:"comment('手机类型（0 安卓 1 ios）') VARCHAR(1)"`
	Phonenum            string    `json:"phonenum" xorm:"comment('设备唯一标识') VARCHAR(100)"`
}

func (c User)TableName() string  {
	return "user"
}

func (c *User) Login(username string,password string) ( user *User)  {
	ok,err := global.TytDB().Where("username=? and password=?",username,password).Get(user)
	if ok{
		return user
	}
	if err != nil {

	}
	return nil
}
