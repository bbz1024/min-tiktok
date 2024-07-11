package variable

import "github.com/streadway/amqp"

const (
	SummeryExchange        = "min-tiktok:summery:exchange"
	SummeryRoutingKey      = "min-tiktok:summery:routingKey"
	SummeryKind            = amqp.ExchangeFanout
	SummeryQueue           = "min-tiktok:summery:queue"
	ExtractSummeryQuestion = "You will be provided with a block of text which is the content of a video, and your task is to give 2 Simplified Chinese sentences to summarize the video."
	ExtractKeyWordQuestion = "You will be provided with a block of text which is the content of a video, and your task is to give 5 tags in Simplified Chinese to the video to attract audience. For example, 美食 | 旅行 | 阅读"
)
