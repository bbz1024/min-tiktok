package mq

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/common/consts/variable"
	"min-tiktok/common/consts/variable/es"
	"min-tiktok/common/util/transform"
	"min-tiktok/services/publish/internal/svc"
)

type ExtractVideoText struct {
	Channel *amqp.Channel
	nsl     *transform.NlsTask
	es      *elastic.Client
	svcCtx  *svc.ServiceContext
}
type ExtractVideoTextReq struct {
	VideoID uint32 `json:"video_id"`
}

var extractText *ExtractVideoText

func InitExtractVideo(svcCtx *svc.ServiceContext, maxPrefetchCnt, consumerCnt int) error {
	extractText = &ExtractVideoText{
		svcCtx: svcCtx,
	}

	// mq
	conn, err := amqp.Dial(svcCtx.Config.RabbitMQ.Dns())
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	extractText.Channel = channel
	if err := channel.Qos(maxPrefetchCnt, 0, false); err != nil {
		return err
	}
	// es
	extractText.es, err = elastic.NewClient(
		elastic.SetURL(svcCtx.Config.Es.Addr),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetGzip(true),
	)
	if err != nil {
		return err
	}
	// declare
	if err := extractText.declare(); err != nil {
		return err
	}

	extractText.nsl = transform.NewNlsTask(
		svcCtx.Config.AlibabaNsl.AccessKeyId,
		svcCtx.Config.AlibabaNsl.AccessKeySecret,
		svcCtx.Config.AlibabaNsl.AppKey,
	)

	for i := 0; i < consumerCnt; i++ {
		name := i
		go extractText.Consumer(name)
	}
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

	// es

	do, err := s.es.IndexExists(es.VideoTextIndex).Do(context.TODO())
	if err != nil {
		return err
	}
	if do {
		return nil
	}
	result, err := s.es.CreateIndex(es.VideoTextIndex).BodyJson(es.VideoTextMapping).Do(context.TODO())
	if err != nil {
		return err
	}
	if !result.Acknowledged {
		return errors.New("es create index failed")
	}
	return nil
}
func (s *ExtractVideoText) Consumer(consumerName int) {
	name := fmt.Sprintf("consumer-%s-%d", variable.ExtractTextQueue, consumerName)
	results, err := s.Channel.Consume(
		variable.ExtractTextQueue,
		name,
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

		// 6. to es
		// TODO ES
		do, err := s.es.Index().
			Index(es.VideoTextIndex).
			Id(fmt.Sprintf("%d", msg.VideoID)).
			BodyJson(map[string]interface{}{
				"video_id":  msg.VideoID,
				"title":     v.Title,
				"user_id":   v.Userid,
				"play_url":  v.Playurl,
				"content":   v.Content.String,
				"create_at": v.CreatedAt.Format("2006-01-02 15:04:05"),
			}).
			Do(context.Background())
		if err != nil {
			logx.Errorf("es error: %v", err)
			continue
		}
		if do.Result != "created" {
			logx.Errorf("es error: %v", err)
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
