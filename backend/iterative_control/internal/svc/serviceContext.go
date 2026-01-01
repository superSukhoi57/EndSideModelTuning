// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"backend/common/llm"
	"iterative_control/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	BLClient *llm.BLClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:   c,
		BLClient: llm.CreateLLMClient(c.LLM.APIKey, c.LLM.Model),
	}
}
