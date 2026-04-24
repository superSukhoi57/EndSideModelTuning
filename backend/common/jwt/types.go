package jwt

import "time"

type Config struct {
	AccessSecret  string
	RefreshSecret string
	AccessExpire  time.Duration
	RefreshExpire time.Duration
	Issuer        string
}
