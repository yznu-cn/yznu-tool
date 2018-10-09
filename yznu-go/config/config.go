package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	ServerConf = &ServerConfig{}
	MySqlConf  = &MySqlConfig{}
	WxConf     = &WxConfig{}
)

type ServerConfig struct {
	Port    string
	RunMode string
}

type MySqlConfig struct {
	Host     string
	Port     string
	DbName   string
	UserName string
	PassWd   string
}

type WxConfig struct {
	AppId  string
	Secret string
}

func init() {
	cfg, err := ini.Load("conf/app.conf")
	if err != nil {
		log.Fatal("load app conf err:", err)
	}
	err = cfg.Section("server").MapTo(ServerConf)
	if err != nil {
		log.Fatal("init server conf err:", err)
	}
	err = cfg.Section("mysql").MapTo(MySqlConf)
	if err != nil {
		log.Fatal("init mysql conf err:", err)
	}
	err = cfg.Section("weixin").MapTo(WxConf)
	if err != nil {
		log.Fatal("init weixin conf err:", err)
	}
}
