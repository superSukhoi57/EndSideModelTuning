package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Redis RedisConf
	MySQL MySQLConf
	LLM   LLMCfg
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

type LLMCfg struct {
	APIKey string
	Model  string
}
