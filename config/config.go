package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	//todo：服务器配置结构
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

var ServConfig AppConfig

//todo:服务器初始化
func InitConfig() *AppConfig {
	file, err := os.Open("./CmsProject/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
