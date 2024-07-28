package mq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"min-tiktok/common/consts/variable"
	client "min-tiktok/common/util/gorse"
	"min-tiktok/common/util/gpt"
	"min-tiktok/models/videoInfo"
	"min-tiktok/services/user/internal/svc"
	"strings"
)

type Deed struct {
	channel        *amqp.Channel
	gorseClient    *client.GorseClient
	videoinfoModel videoInfo.VideoinfoModel
	gpt            *gpt.Gpt
}
type DeedReq struct {
	VideoIds []string `json:"video_id"`
	UserID   string   `json:"user_id"`
}

var Client *Deed

func InitDeed(ctx *svc.ServiceContext) error {
	cli := client.NewGorseClient(ctx.Config.Gorse.GorseAddr, ctx.Config.Gorse.GorseApikey)

	conn, err := amqp.Dial(ctx.Config.RabbitMQ.Dns())
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	Client = &Deed{
		channel:        channel,
		gorseClient:    cli,
		videoinfoModel: videoInfo.NewVideoinfoModel(sqlx.NewMysql(ctx.Config.MySQL.DataSource)),
		gpt:            gpt.NewGpt(ctx.Config.Gpt.ApiKey, ctx.Config.Gpt.ModelID),
	}
	if err := Client.declare(); err != nil {
		return err
	}

	go Client.Consumer()
	return nil
}
func (g *Deed) declare() error {
	if err := g.channel.ExchangeDeclare(variable.DeedExchange, variable.DeedKind,
		true, false, false, false, nil,
	); err != nil {
		return err
	}
	if _, err := g.channel.QueueDeclare(
		variable.DeedQueue, // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	); err != nil {
		return err
	}
	if err := g.channel.QueueBind(
		variable.DeedQueue,      // queue name
		variable.DeedRoutingKey, // routing key
		variable.DeedExchange,   // exchange
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}
func GetInstance() *Deed {
	if Client == nil {
		panic("Client is nil")
	}
	return Client
}
func (g *Deed) Consumer() {
	results, err := g.channel.Consume(
		variable.DeedQueue,
		variable.DeedRoutingKey,
		false, // 关闭自动应答
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	for res := range results {
		var msg DeedReq
		if err := json.Unmarshal(res.Body, &msg); err != nil {
			res.Reject(true)
			logx.Errorw("json unmarshal error: %v", logx.Field("err", err))
			continue
		}

		var category string
		for _, id := range msg.VideoIds {
			c, err := g.videoinfoModel.GetVideoCategory(ctx, id)
			if err != nil {
				res.Reject(true)
				logx.Errorw("get video category error: %v", logx.Field("err", err))
				continue
			}
			category += c
		}
		category, err := g.gpt.ChatWithModel(ctx, variable.DeedQuestion, category)
		if err != nil {
			res.Reject(true)
			logx.Errorw("chat with model error: %v", logx.Field("err", err))
			continue
		}
		// 根据近期种类为用户生成标签
		if _, err := g.gorseClient.UpdateUser(ctx, msg.UserID, client.UserPatch{
			Labels: strings.Split(category, "|"),
		}); err != nil {
			res.Reject(true)
			logx.Errorw("update user error: %v", logx.Field("err", err))
			continue
		}
		res.Ack(false)
	}
}
func (g *Deed) Product(msg DeedReq) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := g.channel.Publish(
		variable.DeedExchange,
		variable.DeedRoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	); err != nil {
		return err
	}
	return nil
}
