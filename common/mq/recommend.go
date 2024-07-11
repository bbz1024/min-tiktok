package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhenghaoz/gorse/client"
	"min-tiktok/common/config"
	"min-tiktok/common/consts/variable"
	"min-tiktok/models/video"
	"time"
)

type GorseRecommend struct {
	channel     *amqp.Channel
	gorseClient *client.GorseClient
	videoModel  video.VideoModel
}
type GorseRecommendReq struct {
	VideoID uint32 `json:"video_id"`
	UserID  uint32 `json:"user_id"`
	Type    variable.Type
}

var gorseClient *GorseRecommend

func InitGorse(gorseConf *config.GorseStructure, MqConf *config.RabbitMQStructure, mysqlConf *config.MysqlStructure) error {
	cli := client.NewGorseClient(gorseConf.GorseAddr, gorseConf.GorseApikey)

	conn, err := amqp.Dial(MqConf.Dns())
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	gorseClient = &GorseRecommend{
		channel:     channel,
		gorseClient: cli,
	}
	if err := gorseClient.declare(); err != nil {
		return err
	}

	mysqlConn := sqlx.NewMysql(mysqlConf.DataSource)
	gorseClient.videoModel = video.NewVideoModel(mysqlConn)
	go gorseClient.Consumer()
	return nil
}
func (g *GorseRecommend) declare() error {
	if err := g.channel.ExchangeDeclare(variable.GorseExchange, variable.GorseKind,
		true, false, false, false, nil,
	); err != nil {
		return err
	}
	if _, err := g.channel.QueueDeclare(
		variable.GorseQueue, // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	); err != nil {
		return err
	}
	if err := g.channel.QueueBind(
		variable.GorseQueue,      // queue name
		variable.GorseRoutingKey, // routing key
		variable.GorseExchange,   // exchange
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}
func GetInstance() *GorseRecommend {
	if gorseClient == nil {
		panic("gorseClient is nil")
	}
	return gorseClient
}
func (g *GorseRecommend) Consumer() {
	results, err := g.channel.Consume(
		variable.GorseQueue,
		variable.GorseRoutingKey,
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
		var err error
		var msg GorseRecommendReq
		if err := json.Unmarshal(res.Body, &msg); err != nil {
			logx.Errorf("unmarshal msg error: %v", err)
			continue
		}
		v, err := g.videoModel.FindOne(ctx, uint64(msg.VideoID))
		if err != nil {
			continue
		}
		now := time.Now().UTC().Format(time.RFC3339)
		switch msg.Type {
		case variable.InsertType:
			_, err = g.gorseClient.InsertItem(ctx, client.Item{
				ItemId:    fmt.Sprintf("%d", v.Id),
				Timestamp: now,
				Comment:   v.Title,
			})
		case variable.FavoriteType:

		}
		if err == nil {
			res.Ack(false)
		}
	}
}
func (g *GorseRecommend) Product(msg GorseRecommendReq) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := g.channel.Publish(
		variable.GorseExchange,
		variable.GorseRoutingKey,
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
