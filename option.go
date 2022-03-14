package pulsar

type (
	option interface {
		apply(options *options)
	}

	options struct {
		label string
	}
)

func defaultOptions() *options {
	return &options{
		label: defaultLabel,
	}
}
