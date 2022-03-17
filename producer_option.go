package pulsar

type (
	producerOption interface {
		applyProducer(options *producerOptions)
	}

	producerOptions struct {
		*options

		serializer serializer
		name       string
		properties map[string]string
	}
)

func defaultProducerOptions() *producerOptions {
	return &producerOptions{
		options: defaultOptions(),

		serializer: serializerUnknown,
		name:       ``,
		properties: make(map[string]string, 0),
	}
}
