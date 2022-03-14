package pulsar

type broker struct {
	// 标签
	Label string `json:"label" yaml:"label" xml:"label" toml:"label" validate:"required"`
	// 连接地址
	// 可以不填，如果不填的话，使用默认地址
	URL string `yaml:"url" json:"url" xml:"url" toml:"url" validate:"required_without=required,hostname_port"`
	// 选项
	Options pulsarOptions `json:"options" yaml:"options" xml:"options" toml:"options"`
}
