package pulsar

import (
	`sync`

	`github.com/apache/pulsar-client-go/pulsar`
)

// Client 客户端
type Client struct {
	clientCache map[string]pulsar.Client

	urlCache     map[string]string
	optionsCache map[string]pulsarOptions

	mutex sync.Mutex
}

func newClient(urls map[string]string, options map[string]pulsarOptions) *Client {
	return &Client{
		clientCache: make(map[string]pulsar.Client, 0),

		urlCache:     urls,
		optionsCache: options,
	}
}

func (c *Client) Producer(topic string, opts ...producerOption) (producer *Producer, err error) {
	_options := defaultProducerOptions()
	for _, opt := range opts {
		opt.applyProducer(_options)
	}

	var client pulsar.Client
	if client, err = c.getPulsar(_options.options); nil != err {
		return
	}

	producer = new(Producer)
	producer.Producer, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic:      topic,
		Name:       _options.name,
		Properties: _options.properties,
	})

	return
}

func (c *Client) Subscribe(topic string, opts ...subscribeOption) (consumer *Consumer, err error) {
	_options := defaultSubscribeOptions()
	for _, opt := range opts {
		opt.applySubscribe(_options)
	}

	var client pulsar.Client
	if client, err = c.getPulsar(_options.options); nil != err {
		return
	}

	consumer = new(Consumer)
	consumer.Consumer, err = client.Subscribe(pulsar.ConsumerOptions{
		Topic:      topic,
		Name:       _options.name,
		Properties: _options.properties,
	})

	return
}

func (c *Client) getPulsar(options *options) (client pulsar.Client, err error) {
	c.mutex.Lock()
	defer func() {
		options.label = defaultLabel
		c.mutex.Unlock()
	}()

	var exist bool
	if client, exist = c.clientCache[options.label]; exist {
		return
	}

	if client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL:               c.urlCache[options.label],
		ConnectionTimeout: c.optionsCache[options.label].Timeout.Connection,
		OperationTimeout:  c.optionsCache[options.label].Timeout.Operation,
		Authentication:    pulsar.NewAuthenticationToken(c.optionsCache[options.label].Token),
	}); nil != err {
		return
	}
	c.clientCache[options.label] = client

	return
}
