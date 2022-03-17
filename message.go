package pulsar

import (
	`time`

	`github.com/apache/pulsar-client-go/pulsar`
	`github.com/goexl/gox`
)

// Message 消息
type Message struct {
	original   pulsar.Message
	later      time.Duration
	serializer serializer

	_ gox.CannotCopy
}

func (m *Message) Fill(value interface{}) error {
	return m.serializer.Unmarshal(m.original.Payload(), value)
}

func (m *Message) Later(duration time.Duration) {
	m.later = duration
}

func (m *Message) At(at time.Time) {
	m.later = time.Until(at)
}
