package pulsar

type (
	producerOption interface {
		applyProducer(options *producerOptions)
	}

	producerOptions struct {
		*options

		name       string
		properties map[string]string
	}
)

func defaultProducerOptions() *producerOptions {
	return &producerOptions{
		options: defaultOptions(),

		name:       ``,
		properties: make(map[string]string, 0),
	}
}
