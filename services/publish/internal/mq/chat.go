package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhenghaoz/gorse/client"
	"min-tiktok/common/consts/variable"
	"min-tiktok/common/util/gpt"
	"min-tiktok/services/publish/internal/svc"
	"strconv"
	"strings"
	"time"
)

var chatText *ChatVideo

type ChatVideo struct {
	Channel     *amqp.Channel
	gpt         *gpt.Gpt
	svcCtx      *svc.ServiceContext
	gorseClient *client.GorseClient
}
type ChatVideoReq struct {
	VideoID uint32 `json:"video_id"`
}

func (v *ChatVideo) declare() error {
	if err := v.Channel.ExchangeDeclare(variable.ChatExchange, variable.ChatKind,
		true, false, false, false, nil,
	); err != nil {
		return err
	}
	if _, err := v.Channel.QueueDeclare(
		variable.ChatQueue, // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	); err != nil {
		return err
	}
	if err := v.Channel.QueueBind(
		variable.ChatQueue,      // queue name
		variable.ChatRoutingKey, // routing key
		variable.ChatExchange,   // exchange
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}
func (v *ChatVideo) Product(msg ChatVideoReq) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := v.Channel.Publish(
		variable.ChatExchange,
		variable.ChatRoutingKey,
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
func (v *ChatVideo) Consumer() {
	results, err := v.Channel.Consume(
		variable.ChatQueue,
		variable.ChatRoutingKey,
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
		var msg ChatVideoReq
		if err := json.Unmarshal(res.Body, &msg); err != nil {
			logx.Errorf("unmarshal msg error: %v", err)
			continue
		}
		logx.Infof("strat consum vidoeID: %d", msg.VideoID)
		vid, err := v.svcCtx.VideoModel.FindOne(context.TODO(), uint64(msg.VideoID))
		if err != nil {
			logx.Errorf("find video error: %v", err)
			continue
		}
		// extract summery
		summery, err := v.gpt.ChatWithModel(context.TODO(), variable.ExtractSummeryQuestion, vid.Content.String)
		if err != nil {
			logx.Errorf("extract summery error: %v", err)
			continue
		}
		// extract keyword
		keywordStr, err := v.gpt.ChatWithModel(context.TODO(), variable.ExtractKeyWordQuestion, vid.Content.String)
		if err != nil {
			logx.Errorf("extract keyword error: %v", err)
			continue
		}
		// extract category
		categoryStr, err := v.gpt.ChatWithModel(context.TODO(), variable.ExtractCategoryQuestion, vid.Content.String)
		if err != nil {
			logx.Errorf("extract category error: %v", err)
			continue
		}
		var conn = sqlx.NewMysql(v.svcCtx.Config.MySQL.DataSource)
		commentTable := "comment"
		videInfoTable := "videoinfo"
		if err := conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
			insertSql := fmt.Sprintf("INSERT INTO %s (videoid,videosummery,keyword,category) VALUES (?,?,?,?)", videInfoTable)
			if _, err := session.ExecCtx(ctx, insertSql,
				msg.VideoID, summery, keywordStr, categoryStr); err != nil {
				return err
			}
			insertSql = fmt.Sprintf("INSERT INTO %s (videoid,userid,content) VALUES (?,?,?)", commentTable)
			if _, err := session.ExecCtx(ctx, insertSql,
				msg.VideoID, vid.Userid, fmt.Sprintf(variable.VideoSummary, summery)); err != nil {
				return err
			}
			if _, err := session.ExecCtx(ctx, insertSql,
				msg.VideoID, vid.Userid, fmt.Sprintf(variable.VideoKeyWord, keywordStr)); err != nil {
				return err
			}
			return nil
		}); err != nil {
			logx.Errorf("transact error: %v", err)
			continue
		}
		// push to gorse
		item := client.Item{
			ItemId:     strconv.Itoa(int(vid.Id)),
			Timestamp:  vid.CreatedAt.UTC().Format(time.RFC3339),
			Comment:    vid.Title,
			Categories: strings.Split(categoryStr, "|"),
			Labels:     strings.Split(keywordStr, "|"),
		}
		row, err := v.gorseClient.InsertItem(context.TODO(), item)
		if err != nil {
			logx.Errorf("insert gorse error: %v", err)
			continue
		}
		if row.RowAffected != 1 {
			logx.Errorf("insert gorse error: %v", err)
			continue
		}
		if err := v.Channel.Ack(res.DeliveryTag, false); err != nil {
			logx.Errorf("ack error: %v", err)
			continue
		}
		logx.Infof("consum vidoeID: %d", msg.VideoID)
	}
}
func GetChatVideo() *ChatVideo {
	return chatText
}
func InitChatVideo(svcCtx *svc.ServiceContext) error {
	chatText = &ChatVideo{
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
	chatText.Channel = channel

	if err := chatText.declare(); err != nil {
		return err
	}
	chatText.gpt = gpt.NewGpt(extractText.svcCtx.Config.Gpt.ApiKey, svcCtx.Config.Gpt.ModelID)
	chatText.gorseClient = client.NewGorseClient(extractText.svcCtx.Config.Gorse.GorseAddr, extractText.svcCtx.Config.Gorse.GorseApikey)
	go chatText.Consumer()
	logx.Infof("chat video init success")
	return nil
}
