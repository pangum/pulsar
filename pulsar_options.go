package pulsar

type pulsarOptions struct {
	// 鉴权
	Token string `yaml:"token" json:"token" xml:"token" toml:"token" validate:"required,jwt"`
	// 超时
	Timeout timeout `yaml:"timeout" json:"timeout" xml:"timeout" toml:"timeout"`
	// 序列化器，默认使用Msgpack做序列化
	Serializer serializer `default:"msgpack" json:"serializer" yaml:"serializer" xml:"serializer" toml:"serializer" validate:"oneof=json msgpack proto xml"`
}
