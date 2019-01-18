package global

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"path/filepath"
	"strings"
)

func InitConfig() {
	ep, err := ExecPath()
	if err != nil {
		panic("获取程序运行文件失败: " + err.Error())
	}
	log.Printf("程序运行文件路径: %s", ep)

	ep = strings.Replace(ep, "\\", "/", -1)
	AppPath = filepath.Dir(ep)
	log.Printf("程序运行文件目录: %s", AppPath)
	confFile := filepath.Join(AppPath,"config" ,Env+".ini")
	fmt.Println(confFile)
	if !IsFile(confFile) {
		panic("获取配置文件失败: " + confFile)
	}
	// 加载配置文件
	Cfg, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: false,
	}, confFile)
	if err != nil {
		panic("配置文件解析失败: " + confFile)
	}
}
