package pulsar

var (
	_ option          = (*optionLabel)(nil)
	_ subscribeOption = (*optionLabel)(nil)
	_ publishOption   = (*optionLabel)(nil)
)

type optionLabel struct {
	label string
}

// Label 配置使用哪一个客户端
func Label(label string) *optionLabel {
	return &optionLabel{
		label: label,
	}
}

// Default 配置使用默认客户端
func Default() *optionLabel {
	return &optionLabel{
		label: defaultLabel,
	}
}

func (l *optionLabel) apply(options *options) {
	options.label = l.label
}

func (l *optionLabel) applySubscribe(options *subscribeOptions) {
	options.label = l.label
}

func (l *optionLabel) applyPublish(options *publishOptions) {
	options.label = l.label
}

func (l *optionLabel) applyBrokers(options *brokersOptions) {
	options.label = l.label
}
