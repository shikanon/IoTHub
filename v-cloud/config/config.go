package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Mysql struct {
	User string
	Pwd  string
	Host string
	Port int `ini:"port"`
	Db   string
}

var MysqlConfig = &Mysql{}

func ConfigInit() {
	// 加载配置文件
	cfg, err := ini.Load("v-cloud/config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	cfg.Section("mysql").MapTo(MysqlConfig)
}
