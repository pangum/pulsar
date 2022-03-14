package pulsar

type config struct {
	// 接入地址
	// 适用于基础使用，只有一个集群
	URL string `yaml:"url" json:"url" xml:"url" toml:"url" validate:"required_without=Brokers,hostname_port"`
	// 服务器列表
	Broker []broker `json:"brokers" yaml:"brokers" xml:"brokers" toml:"brokers" validate:"required_without=URL,dive"`
	// 选项
	Options pulsarOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
