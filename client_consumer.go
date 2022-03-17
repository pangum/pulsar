package pulsar

import (
	`context`
	`fmt`

	`github.com/apache/pulsar-client-go/pulsar`
	`github.com/goexl/exc`
	`github.com/goexl/gox/field`
)

func (c *Client) Handle(ctx context.Context, handler Handler, opts ...subscribeOption) (err error) {
	_options := defaultSubscribeOptions()
	for _, opt := range opts {
		opt.applySubscribe(_options)
	}

	var consumer pulsar.Consumer
	if consumer, err = c.getConsumer(_options); nil != err {
		return
	}

	// 消费消息
	go c.handle(ctx, consumer, handler, c.getSerializer(_options.serializer, _options.options))

	return
}

func (c *Client) handle(ctx context.Context, consumer pulsar.Consumer, handler Handler, serializer serializer) {
	for ; ; {
		message := new(Message)
		if original, err := consumer.Receive(ctx); nil != err {
			break
		} else {
			message.original = original
		}

		// 组装其它属性
		message.serializer = serializer
		if onErr := handler.OnMessage(message); nil != onErr {
			consumer.Nack(message.original)
		} else if 0 != message.later {
			consumer.ReconsumeLater(message.original, message.later)
		} else {
			consumer.Ack(message.original)
		}
	}
}

func (c *Client) getConsumer(options *subscribeOptions) (consumer pulsar.Consumer, err error) {
	c.consumerMutex.Lock()
	defer func() {
		options.key = defaultKey
		c.consumerMutex.Unlock()
	}()

	var client pulsar.Client
	if client, err = c.getPulsar(options.options); nil != err {
		return
	}

	key := fmt.Sprintf(keyFormatter, options.label, options.key)
	var exists bool
	if consumer, exists = c.consumerCache[key]; exists {
		return
	}

	var _topic topic
	if _topic, exists = c.topicCache[key]; !exists {
		err = exc.NewField(`没有找到客户端配置`, field.String(`key`, options.key))
	}
	if nil != err {
		return
	}

	if consumer, err = client.Subscribe(pulsar.ConsumerOptions{
		Topic: _topic.Topic,
		// Name:       topic.name,
		Properties: _topic.Properties,
	}); nil != err {
		return
	}
	c.consumerCache[key] = consumer

	return
}
