package pulsar

var (
	_ = Property

	_ producerOption = (*optionProperty)(nil)
)

type optionProperty struct {
	key   string
	value string
}

// Property 属性
func Property(key string, value string) *optionProperty {
	return &optionProperty{
		key:   key,
		value: value,
	}
}

func (p *optionProperty) applyProducer(options *producerOptions) {
	options.properties[p.key] = p.value
}

func (p *optionProperty) applySubscribe(options *subscribeOptions) {
	options.properties[p.key] = p.value
}
