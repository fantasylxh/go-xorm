package tyt

import (
	"time"
)

type UserLine struct {
	Id             int       `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	UserIp         string    `json:"user_ip" xorm:"not null comment('åŸŸå') VARCHAR(30)"`
	Area           string    `json:"area" xorm:"comment('åœ°åŒº') VARCHAR(255)"`
	IsDel          string    `json:"is_del" xorm:"default '0' comment('0 ä¸åˆ é™¤ 1  åˆ é™¤') CHAR(1)"`
	CreateDate     time.Time `json:"create_date" xorm:"DATETIME"`
	Port           string    `json:"port" xorm:"comment('ç«¯å£') VARCHAR(10)"`
	Http           string    `json:"http" xorm:"comment('åè®®') VARCHAR(10)"`
	Level          int       `json:"level" xorm:"INT(11)"`
	DefCount       int       `json:"def_count" xorm:"default 00000000000 INT(11)"`
	RelCount       int       `json:"rel_count" xorm:"default 1 comment('Ã¥Â®Å¾Ã©â„¢â€¦Ã©â€¡Â') INT(11)"`
	ComCount       int       `json:"com_count" xorm:"comment('æ¯”è¾ƒ1') INT(11)"`
	ComCountTwo    int       `json:"com_count_two" xorm:"comment('æ¯”è¾ƒ2') INT(11)"`
	Status         string    `json:"status" xorm:"comment('çŠ¶æ€ 0 æœªæ¿€æ´» 1 å·²ç»æ¿€æ´»') CHAR(1)"`
	LineIp         string    `json:"line_ip" xorm:"comment('åŸŸåæ‰€å¯¹åº”çš„ipåœ°å€') VARCHAR(30)"`
	IsFree         int       `json:"is_free" xorm:"default 1 comment('æ˜¯å¦å…è´¹çº¿è·¯') INT(1)"`
	Html           string    `json:"html" xorm:"TEXT"`
	IpCount        int       `json:"ip_count" xorm:"default 0 INT(11)"`
	Type           string    `json:"type" xorm:"VARCHAR(255)"`
	EncryptionMode string    `json:"encryption_mode" xorm:"VARCHAR(255)"`
	Password       string    `json:"password" xorm:"VARCHAR(255)"`
	ImgUrl         string    `json:"img_url" xorm:"VARCHAR(255)"`
	Pid            int       `json:"pid" xorm:"INT(10)"`
}
