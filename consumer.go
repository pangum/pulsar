package pulsar

import (
	`context`

	`github.com/apache/pulsar-client-go/pulsar`
	`github.com/goexl/gox`
)

// Consumer 消费者封装
type Consumer struct {
	consumer   pulsar.Consumer
	serializer serializer

	_ gox.CannotCopy
}

func (c *Consumer) Handle(ctx context.Context, handler Handler) (err error) {
	for ; ; {
		message := new(Message)
		if message.original, err = c.consumer.Receive(ctx); nil != err {
			break
		}

		// 组装其它属性
		message.serializer = c.serializer
		if onErr := handler.OnMessage(message); nil != onErr {
			c.consumer.Nack(message.original)
		} else if 0 != message.later {
			c.consumer.ReconsumeLater(message.original, message.later)
		} else {
			c.consumer.Ack(message.original)
		}
	}

	return
}

func (c *Consumer) Close() {
	c.consumer.Close()
}
