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

type Mongodb struct {
	Host                string
	Port                int `ini:"port"`
	Db                  string
	BaseModelConcise    string `ini:"concise"`
	BaseModelIntact     string `ini:"intact"`
	IntactProductModel  string `ini:"intact_model"`
	ConciseProductModel string `ini:"concise_model"`
}

var MongodbConfig = &Mongodb{}

type Constant struct {
	TimeFormat           string
	ModelNumber          int `ini:"modelNumber"`
	NodeTypeNumber       int `ini:"nodeTypeNumber"`
	NetWorkNumber        int `ini:"netWorkNumber"`
	DataFormatNumber     int `ini:"dataFormatNumber"`
	AuthMethodNumber     int `ini:"authMethodNumber"`
	TopicOperationNumber int `ini:"topicOperationNumber"`
}

var GeneralConfig = &Constant{}

func ConfigInit() {
	// 加载配置文件
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	cfg.Section("mysql").MapTo(MysqlConfig)
	cfg.Section("mongodb").MapTo(MongodbConfig)
	cfg.Section("constant").MapTo(GeneralConfig)
}
