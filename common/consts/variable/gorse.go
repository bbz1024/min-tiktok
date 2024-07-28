package variable

import "github.com/streadway/amqp"

const (
	GorseExchange   = "min-tiktok:gorse:exchange"
	GorseRoutingKey = "min-tiktok:gorse:routingKey"
	GorseKind       = amqp.ExchangeFanout
	GorseQueue      = "min-tiktok:gorse:queue"
)
