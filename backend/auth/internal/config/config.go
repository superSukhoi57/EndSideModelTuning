package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	FeishuAuth FeishuAuthCfg
}

type FeishuAuthCfg struct {
	AppID       string
	AppSecret   string
	RedirectURI string
}
