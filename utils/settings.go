package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	AppPort string

	JwtKey string

	Db         string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置参数有误，请检查后重试！")
		return
	}
	LoadServer(cfg)
	LoadDatabase(cfg)
}
func LoadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").String()
	AppPort = cfg.Section("server").Key("AppPort").String()
	JwtKey = cfg.Section("server").Key("JwtKey").String()
}

func LoadDatabase(cfg *ini.File) {
	Db = cfg.Section("database").Key("Db").String()
	DbUser = cfg.Section("database").Key("DbUser").String()
	DbPassword = cfg.Section("database").Key("DbPassword").String()
	DbHost = cfg.Section("database").Key("DbHost").String()
	DbPort = cfg.Section("database").Key("DbPort").String()
	DbName = cfg.Section("database").Key("DbName").String()
}
