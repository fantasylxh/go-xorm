package main

import (
	"gameCityService/echo_web"
	"gameCityService/global"
)

var (
	ENVIRONMENT = "developer" // 全局环境
)
func main() {
	if ENVIRONMENT == global.Dev {
		global.Env = global.Dev
	} else if ENVIRONMENT == global.Test {
		global.Env = global.Test
	} else {
		global.Env = global.Prod
	}
	global.InitConfig();
	global.InitDB([]string{
		global.ADVERT,
		global.TYT,
	})
	global.InitLog([]string{global.AppCron})
	global.InitRedis()
	echo_web.EchoInit()

}