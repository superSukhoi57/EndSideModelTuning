package svc

import (
	"fmt"

	"backend/common/llm"
	"backend/common/protocol/authpb"
	"iterative_control/internal/config"
	"iterative_control/internal/dao/machine"
	"iterative_control/internal/dao/parameter"
	"iterative_control/internal/dao/result"
	"iterative_control/internal/dao/task"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config       config.Config
	Rdb          *redis.Client
	DB           *gorm.DB
	BLClient     *llm.BLClient
	AuthClient   authpb.AuthServiceClient
	MachineDAO   *machine.MachineDAO
	ParameterDAO *parameter.ParameterDAO
	TaskDAO      *task.TaskDAO
	ResultDAO    *result.ResultDAO
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

	conn, err := grpc.Dial(c.Auth.GRPCAddr, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("failed to connect auth grpc server: %v", err))
	}

	return &ServiceContext{
		Config:       c,
		Rdb:          rdb,
		DB:           db,
		BLClient:     llm.CreateLLMClient(c.LLM.APIKey, c.LLM.Model),
		AuthClient:   authpb.NewAuthServiceClient(conn),
		MachineDAO:   machine.NewMachineDAO(db),
		ParameterDAO: parameter.NewParameterDAO(db),
		TaskDAO:      task.NewTaskDAO(db),
		ResultDAO:    result.NewResultDAO(db),
	}
}
