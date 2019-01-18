package tyt

import (
	"time"
)

type SysUsers struct {
	Id               int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20)"`
	Username         string    `json:"username" xorm:"comment('用户名') VARCHAR(20)"`
	Mobile           string    `json:"mobile" xorm:"comment('手机号码') VARCHAR(20)"`
	Password         string    `json:"password" xorm:"comment('密码') VARCHAR(20)"`
	Name             string    `json:"name" xorm:"comment('姓名') VARCHAR(50)"`
	Post             string    `json:"post" xorm:"comment('岗位') VARCHAR(20)"`
	Token            string    `json:"token" xorm:"comment('登录token') VARCHAR(100)"`
	BusName          string    `json:"bus_name" xorm:"comment('商户名称') VARCHAR(100)"`
	BusAddr          string    `json:"bus_addr" xorm:"comment('商户地址') VARCHAR(200)"`
	Idcard           string    `json:"idcard" xorm:"comment('身份证号') VARCHAR(50)"`
	InvitationCode   string    `json:"invitation_code" xorm:"comment('邀请码') VARCHAR(100)"`
	ExtensionWorkers string    `json:"extension_workers" xorm:"comment('推广人员') VARCHAR(20)"`
	CateLevel1       int64     `json:"cate_level1" xorm:"comment('商户经营的品类一级') BIGINT(20)"`
	CateLevel2       int64     `json:"cate_level2" xorm:"comment('商户经营的品类二级') BIGINT(20)"`
	AccountType      int       `json:"account_type" xorm:"comment('用户类型(0后台管理账号1商户2司机3外勤人员可app/登陆后台管理)') INT(2)"`
	LastLoginTime    time.Time `json:"last_login_time" xorm:"comment('上次登录时间') DATETIME"`
	CreateTime       time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	DelFlg           string    `json:"del_flg" xorm:"default 'N' comment('逻辑删除(N启用F禁用)') CHAR(1)"`
	CreateUser       int64     `json:"create_user" xorm:"comment('创建人') BIGINT(20)"`
	UpdateUser       int64     `json:"update_user" xorm:"comment('修改人') BIGINT(20)"`
	UpdateTime       time.Time `json:"update_time" xorm:"comment('修改时间') DATETIME"`
	UserType         string    `json:"user_type" xorm:"comment('客户类型') VARCHAR(100)"`
}
