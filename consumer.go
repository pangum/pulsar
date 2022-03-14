package pulsar

import (
	`github.com/apache/pulsar-client-go/pulsar`
)

// Consumer 消费者封装
type Consumer struct {
	pulsar.Consumer
}
