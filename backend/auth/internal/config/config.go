package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	FeishuAuth FeishuAuthCfg
	Redis      RedisConf
}

type FeishuAuthCfg struct {
	AppID       string
	AppSecret   string
	RedirectURI string
}

type RedisConf struct {
	Host string
	Pass string
	DB   int
}
