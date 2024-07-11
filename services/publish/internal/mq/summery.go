package mq

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhenghaoz/gorse/client"
	"min-tiktok/common/consts/variable"
	mq2 "min-tiktok/common/mq"
	"min-tiktok/common/util/gpt"
	"min-tiktok/common/util/transform"
	"min-tiktok/services/publish/internal/svc"
)

type VideoSummery struct {
	Channel     *amqp.Channel
	nsl         *transform.NlsTask
	gpt         *gpt.Gpt
	svcCtx      *svc.ServiceContext
	gorseClient *client.GorseClient
}
type VideoSummeryReq struct {
	VideoID uint32 `json:"video_id"`
}

var summery *VideoSummery

func InitVideoSummery(svcCtx *svc.ServiceContext) error {
	summery = &VideoSummery{
		svcCtx: svcCtx,
	}
	conn, err := amqp.Dial(svcCtx.Config.RabbitMQ.Dns())
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	summery.Channel = channel

	if err := summery.declare(); err != nil {
		return err
	}

	summery.nsl = transform.NewNlsTask(
		svcCtx.Config.AlibabaNsl.AccessKeyId,
		svcCtx.Config.AlibabaNsl.AccessKeySecret,
		svcCtx.Config.AlibabaNsl.AppKey,
	)
	summery.gpt = gpt.NewGpt(summery.svcCtx.Config.Gpt.ApiKey, svcCtx.Config.Gpt.ModelID)
	summery.gorseClient = client.NewGorseClient(summery.svcCtx.Config.Gorse.GorseAddr, summery.svcCtx.Config.Gorse.GorseApikey)
	go summery.Consumer()
	return nil
}
func GetInstance() *VideoSummery {
	if summery == nil {
		panic("summery is nil")
	}
	return summery
}
func (s *VideoSummery) declare() error {
	if err := s.Channel.ExchangeDeclare(variable.SummeryExchange, variable.SummeryKind,
		true, false, false, false, nil,
	); err != nil {
		return err
	}
	if _, err := s.Channel.QueueDeclare(
		variable.SummeryQueue, // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	); err != nil {
		return err
	}
	if err := s.Channel.QueueBind(
		variable.SummeryQueue,      // queue name
		variable.SummeryRoutingKey, // routing key
		variable.SummeryExchange,   // exchange
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}
func (s *VideoSummery) Consumer() {
	results, err := s.Channel.Consume(
		variable.SummeryQueue,
		variable.SummeryRoutingKey,
		false, // 关闭自动应答
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	for res := range results {
		var msg VideoSummeryReq
		if err := json.Unmarshal(res.Body, &msg); err != nil {
			logx.Errorf("unmarshal msg error: %v", err)
			continue
		}
		logx.Infof("strat consum vidoeID: %d", msg.VideoID)
		// 1. get video info by db
		v, err := s.svcCtx.VideoModel.FindOne(context.Background(), uint64(msg.VideoID))
		if err != nil {
			logx.Errorf("get video info error: %v", err)
			continue
		}
		// 2. get video text
		taskID, err := s.nsl.SubmitTask(v.Playurl)
		if err != nil {
			logx.Errorf("get video text error: %v", err)
			continue
		}
		result, err := s.nsl.GetTaskResult(taskID)
		if err != nil {
			logx.Errorf("get video text error: %v", err)
			continue
		}
		// 3. extract summery
		response, err := s.gpt.ChatWithModel(context.Background(), variable.ExtractSummeryQuestion, result.Content)
		if err != nil {
			logx.Errorf("extract summery error: %v", err)
			continue
		}
		// 4. update db
		v.Summery = sql.NullString{String: response, Valid: true}
		if err = s.svcCtx.VideoModel.Update(context.Background(), v); err != nil {
			logx.Errorf("update db error: %v", err)
			continue
		}
		//5. insert gorse to make recommend
		if err := mq2.GetInstance().Product(
			mq2.GorseRecommendReq{VideoID: uint32(v.Id)},
		); err != nil {
			return
		}
		if err = s.Channel.Ack(res.DeliveryTag, false); err != nil {
			logx.Errorf("ack error: %v", err)
			continue
		}
		logx.Infof("consum vidoeID: %d success", msg.VideoID)
	}
}
func (s *VideoSummery) Product(msg VideoSummeryReq) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := s.Channel.Publish(
		variable.SummeryExchange,
		variable.SummeryRoutingKey,
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
