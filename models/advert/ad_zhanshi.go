package advert

import (
	"fmt"
	"gameCityService/global"
	"github.com/satori/go.uuid"
	"time"

	_ "github.com/satori/go.uuid"
)

type AdZhanshi struct {
	Id         string `json:"id" xorm:"not null pk CHAR(36)"`
	AdId       int    `json:"ad_id" xorm:"not null comment('广告id') INT(11)"`
	Device     int    `json:"device" xorm:"not null default 1 comment('1:android,2:iOS') SMALLINT(2)"`
	CreateTime int64    `json:"create_time" xorm:"not null comment('展示时间') INT(11)"`
	IsClick    int    `json:"is_click" xorm:"not null default 0 comment('点击') TINYINT(1)"`
}

func (AdZhanshi) TableName() string {
	return global.TablePrefixGameCity + "zhanshi";
}

func (t AdZhanshi)CreateData(ad_id int,device int) string  {
	id := uuid.NewV4().String()

	res,err := global.AdvertDB().Table(t.TableName()).InsertOne(AdZhanshi{Id:id,AdId:ad_id,Device:device,CreateTime:time.Now().Unix()})
	if res > 0 && err == nil{
		return id
	}
	return ""
}

func (t AdZhanshi) AdvertClick(zs_id string)  {
	engine := global.AdvertDB()
	res,err := engine.Table(t.TableName()).ID(zs_id).Update(&AdZhanshi{IsClick:1})

	fmt.Println(engine.LastSQL())
	fmt.Println(err)
	new(global.Logger).Trace("%i",res)
}
