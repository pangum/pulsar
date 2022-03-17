package pulsar

type (
	subscribeOption interface {
		applySubscribe(options *subscribeOptions)
	}

	subscribeOptions struct {
		*options

		key        string
		serializer serializer
		name       string
		properties map[string]string
	}
)

func defaultSubscribeOptions() *subscribeOptions {
	return &subscribeOptions{
		options: defaultOptions(),

		key:        defaultKey,
		serializer: serializerUnknown,
		name:       ``,
		properties: make(map[string]string, 0),
	}
}
