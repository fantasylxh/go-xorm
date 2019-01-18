package global

import (
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"gopkg.in/ini.v1"
	"time"
)

// ----常量----
const (
	Dev  = "development"
	Test = "test"
	Prod = "production"
	LIMIT = 50
)

const (
	HTTP_OK = 0
	HTTP_TOKEN_FX = 50008 //非法的token
	HTTP_TOKEN_LOGOUT = 50012 //其他客户端登录了
	HTTP_TOKEN_GQ = 50014 //Token 过期了
	HTTP_ERR = 40000
)

const (
	Windows = "windows"
	Linux   = "linux"
	Darwin  = "darwin"
)
const (
	AppCron = "cron"
)

const (
	ADVERT          = "advert"
	TYT    			= "tyt"
)

const (
	TablePrefixGameCity = "ad_"
)


const (
	TimeParse         = "2006-01-02 15:04:05"
	TimeSortParse     = "2006-01-02"
	TimeSortLinkParse = "20060102"
)




// ----变量----
var (
	Env       string                  // 环境
	AppPath   string                  // 程序运行目录
	Cfg       *ini.File               // 配置
	orm       map[string]*xorm.Engine // 数据库engine
	RedisConn *redis.Client           // redis连接
	Logs      map[string]*Logger      // log
)

var (
	Location = time.FixedZone("CST", 8*3600)
)
