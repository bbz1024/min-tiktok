package variable

import "github.com/streadway/amqp"

const (
	ExtractTextExchange     = "min-tiktok:text:exchange"
	ExtractTextRoutingKey   = "min-tiktok:text:routingKey"
	ExtractTextKind         = amqp.ExchangeFanout
	ExtractTextQueue        = "min-tiktok:text:queue"
	ExtractSummeryQuestion  = "You will be provided with a block of text which is the content of a video, and your task is to give 2 Simplified Chinese sentences to summarize the video."
	ExtractKeyWordQuestion  = "You will be provided with a block of text which is the content of a video, and your task is to give 5 keyword in Simplified Chinese to the video to attract audience. For example, 美食 | 旅行 | 阅读"
	ExtractCategoryQuestion = "You will be provided with a block of text which is the content of a video, and your task is to give 5 category in Simplified Chinese to the video to attract audience. For example, 游戏 | 电影 | 旅游 | 解说"
	VideoSummary            = "我是视频助手，该视频的摘要是：%s"
	VideoKeyWord            = "我是视频助手，该视频的关键词是：%s"
)
