package advert

import "gameCityService/global"

type AdBanner struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Title      string `json:"title" xorm:"not null comment('广告标题') VARCHAR(128)"`
	Url        string `json:"url" xorm:"not null comment('广告主链接') VARCHAR(255)"`
	ImgUrl     string `json:"img_url" xorm:"not null comment('广告图片') CHAR(255)"`
	Status     int    `json:"status" xorm:"not null default 1 comment('状态') TINYINT(1)"`
	CreateTime int    `json:"create_time" xorm:"not null comment('创建时间') INT(11)"`
	UpdateTime int    `json:"update_time" xorm:"not null comment('最后一尺修改时间') INT(11)"`
	CountDown int 	  `json:"count_down" xorm:"not null comment('广告显示时间') INT(11)" `
}

func (c AdBanner) TableName() string  {
	return global.TablePrefixGameCity + "banner"
}

func (c *AdBanner) GetBanner() *AdBanner  {
	banner := new(AdBanner)
	ok,err:= global.AdvertDB().Table(c.TableName()).Get(banner)
	if ok && err == nil{
		return banner
	}
	return nil
}