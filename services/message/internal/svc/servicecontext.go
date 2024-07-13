package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/models/message"
	"min-tiktok/services/message/internal/config"
	"min-tiktok/services/relation/relationclient"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel message.MessagesModel
	RelationRpc  relationclient.Relation
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:       c,
		MessageModel: message.NewMessagesModel(mysqlConn),
		RelationRpc:  relationclient.NewRelation(zrpc.MustNewClient(c.RelationRpc)),
	}
}
