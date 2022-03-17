package pulsar

import (
	`github.com/apache/pulsar-client-go/pulsar`
	`github.com/goexl/gox`
)

// Producer 生产者
type Producer struct {
	producer   pulsar.Producer
	serializer serializer
	_          gox.CannotCopy
}

func (p *Producer) Send(payload interface{}) (err error) {
	return
}

func (p *Producer) Close() {
	p.producer.Close()
}
