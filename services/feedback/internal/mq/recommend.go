package recommend

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhenghaoz/gorse/client"
	"min-tiktok/common/consts/variable"
	"min-tiktok/services/feedback/internal/svc"
	"strconv"
	"time"
)

type GorseRecommend struct {
	channel     *amqp.Channel
	gorseClient *client.GorseClient
}
type GorseRecommendReq struct {
	VideoIds []uint32 `json:"video_id"`
	UserID   uint32   `json:"user_id"`
	Type     variable.FeedType
}

var gorseClient *GorseRecommend

func InitGorse(ctx *svc.ServiceContext) error {
	cli := client.NewGorseClient(ctx.Config.Gorse.GorseAddr, ctx.Config.Gorse.GorseApikey)

	conn, err := amqp.Dial(ctx.Config.RabbitMQ.Dns())
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
		var msg GorseRecommendReq
		if err := json.Unmarshal(res.Body, &msg); err != nil {
			logx.Errorf("unmarshal msg error: %v", err)
			continue
		}
		now := time.Now().UTC().Format(time.RFC3339)
		var feedback []client.Feedback
		for _, videoid := range msg.VideoIds {
			feedback = append(feedback, client.Feedback{
				FeedbackType: string(msg.Type),
				UserId:       strconv.Itoa(int(msg.UserID)),
				ItemId:       strconv.Itoa(int(videoid)),
				Timestamp:    now,
			})
		}
		if _, err := g.gorseClient.InsertFeedback(ctx, feedback); err != nil {
			logx.Errorf("insert feedback error: %v", err)
			continue
		}
		if err := res.Ack(false); err != nil {
			logx.Errorf("ack error: %v", err)
			continue
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
