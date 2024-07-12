package mq

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/common/consts/variable"
	"min-tiktok/common/util/transform"
	"min-tiktok/services/publish/internal/svc"
)

type ExtractVideoText struct {
	Channel *amqp.Channel
	nsl     *transform.NlsTask

	svcCtx *svc.ServiceContext
}
type ExtractVideoTextReq struct {
	VideoID uint32 `json:"video_id"`
}

var extractText *ExtractVideoText

func InitExtractVideo(svcCtx *svc.ServiceContext) error {
	extractText = &ExtractVideoText{
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
	extractText.Channel = channel

	if err := extractText.declare(); err != nil {
		return err
	}

	extractText.nsl = transform.NewNlsTask(
		svcCtx.Config.AlibabaNsl.AccessKeyId,
		svcCtx.Config.AlibabaNsl.AccessKeySecret,
		svcCtx.Config.AlibabaNsl.AppKey,
	)
	go extractText.Consumer()
	logx.Infof("extract text  init success")
	return nil
}
func GetExtractVideoText() *ExtractVideoText {
	if extractText == nil {
		panic("extractText is nil")
	}
	return extractText
}
func (s *ExtractVideoText) declare() error {
	if err := s.Channel.ExchangeDeclare(variable.ExtractTextExchange, variable.ExtractTextKind,
		true, false, false, false, nil,
	); err != nil {
		return err
	}
	if _, err := s.Channel.QueueDeclare(
		variable.ExtractTextQueue, // name
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	); err != nil {
		return err
	}
	if err := s.Channel.QueueBind(
		variable.ExtractTextQueue,      // queue name
		variable.ExtractTextRoutingKey, // routing key
		variable.ExtractTextExchange,   // exchange
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}
func (s *ExtractVideoText) Consumer() {
	results, err := s.Channel.Consume(
		variable.ExtractTextQueue,
		variable.ExtractTextRoutingKey,
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
		var msg ExtractVideoTextReq
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
		// 4. update db
		v.Content = sql.NullString{String: result.Content, Valid: true}
		if err = s.svcCtx.VideoModel.Update(context.Background(), v); err != nil {
			logx.Errorf("update db error: %v", err)
			continue
		}
		// 5. extract task
		if err := GetChatVideo().Product(ChatVideoReq{VideoID: msg.VideoID}); err != nil {
			logx.Errorf("extract task error: %v", err)
			continue
		}
		if err = s.Channel.Ack(res.DeliveryTag, false); err != nil {
			logx.Errorf("ack error: %v", err)
			continue
		}
		logx.Infof("consum vidoeID: %d success", msg.VideoID)
	}
}
func (s *ExtractVideoText) Product(msg ExtractVideoTextReq) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := s.Channel.Publish(
		variable.ExtractTextExchange,
		variable.ExtractTextRoutingKey,
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
