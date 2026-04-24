package svc

import (
	"fmt"

	"auth/internal/config"
	"auth/internal/dao"

	commonjwt "backend/common/jwt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config   config.Config
	Rdb      *redis.Client
	DB       *gorm.DB
	StateDAO *dao.StateDAO
	UserDAO  *dao.UserDAO
	TokenDAO *dao.TokenDAO
	JWTMgr   *commonjwt.Manager
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Pass,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.DB,
		c.MySQL.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	jwtCfg := commonjwt.Config{
		AccessSecret:  c.JWT.AccessSecret,
		RefreshSecret: c.JWT.RefreshSecret,
		AccessExpire:  c.JWT.AccessExpire,
		RefreshExpire: c.JWT.RefreshExpire,
		Issuer:        c.JWT.Issuer,
	}
	jwtMgr := commonjwt.NewManager(jwtCfg)

	return &ServiceContext{
		Config:   c,
		Rdb:      rdb,
		DB:       db,
		StateDAO: dao.NewStateDAO(rdb),
		UserDAO:  dao.NewUserDAO(db),
		TokenDAO: dao.NewTokenDAO(rdb),
		JWTMgr:   jwtMgr,
	}
}
