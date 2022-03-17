package pulsar

// Handler 消费器
type Handler interface {
	// OnMessage 当有消息时
	OnMessage(message *Message) (err error)
}
