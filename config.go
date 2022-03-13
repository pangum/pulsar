package pulsar

type config struct {
	// Apache组织
	Apache struct {
		// 主体配置
		Pulsar pulsarConfig `yaml:"pulsar" json:"pulsar" xml:"pulsar" toml:"pulsar"`
	} `yaml:"apache" json:"apache" xml:"apache" toml:"apache"`
}
