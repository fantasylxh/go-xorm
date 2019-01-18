package tyt

type ActIdUser struct {
	Id        string `json:"ID_" xorm:"not null pk default '' VARCHAR(64)"`
	Rev       int    `json:"REV_" xorm:"INT(11)"`
	First     string `json:"FIRST_" xorm:"VARCHAR(255)"`
	Last      string `json:"LAST_" xorm:"VARCHAR(255)"`
	Email     string `json:"EMAIL_" xorm:"VARCHAR(255)"`
	Pwd       string `json:"PWD_" xorm:"VARCHAR(255)"`
	PictureId string `json:"PICTURE_ID_" xorm:"VARCHAR(64)"`
}
