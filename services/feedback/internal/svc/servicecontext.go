package svc

import (
	"github.com/zhenghaoz/gorse/client"
	"min-tiktok/services/feedback/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Recommend *client.GorseClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Recommend: client.NewGorseClient(c.Gorse.GorseAddr, c.Gorse.GorseApikey),
	}
}
