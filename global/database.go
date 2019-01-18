package global

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "gopkg.in/clog.v1"
	"path"
)

func InitDB(dbNames []string) {
	if len(dbNames) == 0 {
		panic("没有传入数据库名称")
	}
	orm = make(map[string]*xorm.Engine, len(dbNames))
	for _, dbName := range dbNames {
		db := Cfg.Section("db." + dbName)
		host := db.Key("HOST").MustString("127.0.0.1")
		port := db.Key("PORT").MustString("3306")
		name := db.Key("NAME").MustString("advert")
		user := db.Key("USER").MustString("root")
		pass := db.Key("PASS").MustString("Root2018@")
		charset := db.Key("CHARSET").MustString("utf8")
		engine, err := xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name+"?charset="+charset+"&interpolateParams=true")
		if err != nil {
			panic("数据库engine创建失败: " + err.Error())
		}
		err = engine.Ping()
		if err != nil {
			panic("数据库连接失败: " + err.Error())
		}
		engine.SetMapper(core.GonicMapper{})
		l := Cfg.Section("log")
		xl, err := log.NewFileWriter(path.Join(path.Join(AppPath, "log"), dbName+".db.log"), log.FileRotationConfig{
			Rotate:  l.Key("ROTATE").MustBool(true),
			Daily:   l.Key("DAILY").MustBool(true),
			MaxSize: 1 << uint(l.Key("MAX_SIZE_SHIFT").MustInt(28)),
			MaxDays: l.Key("MAX_DAYS").MustInt64(30),
		})
		if err != nil {
			log.Fatal(2, "数据库日志文件创建失败: %v", err)
		}
		engine.SetLogger(xorm.NewSimpleLogger3(xl, xorm.DEFAULT_LOG_PREFIX, xorm.DEFAULT_LOG_FLAG, core.LOG_DEBUG))
		engine.ShowSQL(true)
		orm[dbName] = engine
	}
}

func AdvertDB() *xorm.Session {
	s := orm[ADVERT].NewSession()
	defer s.Close()
	return s
}

func TytDB() *xorm.Session {
	s := orm[TYT].NewSession()
	defer s.Close()
	return s
}

