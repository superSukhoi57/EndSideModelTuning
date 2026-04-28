package config

import (
	"time"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	FeishuAuth FeishuAuthCfg
	Redis      RedisConf
	MySQL      MySQLConf
	JWT        JWTConf
	Frontend   FrontendConf
	GRPC       GRPCConf
}

type GRPCConf struct {
	Port int
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

type MySQLConf struct {
	Host    string
	Port    int
	User    string
	Pass    string
	DB      string
	Charset string
}

type JWTConf struct {
	AccessSecret  string
	RefreshSecret string
	AccessExpire  time.Duration
	RefreshExpire time.Duration
	Issuer        string
}

type FrontendConf struct {
	RedirectURI string
}
