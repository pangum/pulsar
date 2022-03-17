package pulsar

import (
	`sync`

	`github.com/apache/pulsar-client-go/pulsar`
)

// Client 客户端
type Client struct {
	clientCache   map[string]pulsar.Client
	consumerCache map[string]pulsar.Consumer

	urlCache     map[string]string
	optionsCache map[string]pulsarOptions
	topicCache   map[string]topic

	clientMutex   sync.Mutex
	consumerMutex sync.Mutex
}

func newClient(urls map[string]string, options map[string]pulsarOptions, topics map[string]topic) *Client {
	return &Client{
		clientCache: make(map[string]pulsar.Client, 0),

		urlCache:     urls,
		optionsCache: options,
		topicCache:   topics,
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
	producer.serializer = c.getSerializer(_options.serializer, _options.options)
	producer.producer, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic:      topic,
		Name:       _options.name,
		Properties: _options.properties,
	})

	return
}

func (c *Client) getPulsar(options *options) (client pulsar.Client, err error) {
	c.clientMutex.Lock()
	defer func() {
		options.label = defaultLabel
		c.clientMutex.Unlock()
	}()

	var exists bool
	if client, exists = c.clientCache[options.label]; exists {
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

func (c *Client) getSerializer(serializer serializer, options *options) (final serializer) {
	if serializerUnknown == serializer {
		final = c.optionsCache[options.label].Serializer
	} else {
		final = serializer
	}

	return
}
