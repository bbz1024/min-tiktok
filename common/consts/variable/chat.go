package variable

import "github.com/streadway/amqp"

const (
	ChatExchange   = "min-tiktok:chat:exchange"
	ChatRoutingKey = "min-tiktok:chat:routingKey"
	ChatKind       = amqp.ExchangeFanout
	ChatQueue      = "min-tiktok:chat:queue"
)

type FeedType string

const (
	FavoriteFeedBack FeedType = "favorite"
	CommentFeedBack           = "comment"
	ReadFeedBack              = "read"
)
