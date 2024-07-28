package variable

import "github.com/streadway/amqp"

const (
	DeedExchange   = "min-tiktok:deed:exchange"
	DeedRoutingKey = "min-tiktok:deed:routingKey"
	DeedKind       = amqp.ExchangeFanout
	DeedQueue      = "min-tiktok:deed:queue"
	DeedQuestion   = "You are given a piece of text about the categories of videos the user has recently watched and your task is to extract 3 or 4 categories the user likes in Simplified Chinese to engage the audience." +
		" For example: game | movie | travel | commentary"
)
