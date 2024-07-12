package variable

import "github.com/streadway/amqp"

type Type string

const (
	GorseExchange   = "min-tiktok:gorse:exchange"
	GorseRoutingKey = "min-tiktok:gorse:routingKey"
	GorseKind       = amqp.ExchangeFanout
	GorseQueue      = "min-tiktok:gorse:queue"
)
