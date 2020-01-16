package main

import (
	"github.com/shikanon/IoTOrbHub/config"
	"github.com/shikanon/IoTOrbHub/pkg/apiserver"
)

func init() {
	// 初始化配置
	config.ConfigInit()
}

func main() {
	apiserver.ApiRegister()
}
