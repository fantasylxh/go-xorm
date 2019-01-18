package advert

import "gameCityService/global"

type AdAdmin struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username     string `json:"username" xorm:"not null CHAR(16)"`
	Password     string `json:"password" xorm:"not null CHAR(64)"`
	Avatar       string `json:"avatar" xorm:"not null default '' VARCHAR(128)"`
	Introduction string `json:"introduction" xorm:"not null default '' VARCHAR(128)"`
	CreateTime   int    `json:"create_time" xorm:"not null INT(11)"`
	UpdateTime   int    `json:"update_time" xorm:"not null INT(11)"`
	Status       int    `json:"status" xorm:"not null default 1 TINYINT(1)"`
}

func (c AdAdmin)TableName() string  {
	return global.TablePrefixGameCity + "admin"
}
func (c * AdAdmin)GetAdminByUserName(username string) ( *AdAdmin,  error)  {
	admin:= new(AdAdmin)
	ok,err :=  global.AdvertDB().Table(c.TableName()).Where("username=?",username).Get(admin)
	if ok && err != nil{
		return admin,nil
	}
	return nil,err
}