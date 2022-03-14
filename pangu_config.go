package pulsar

type panguConfig struct {
	// Apache组织
	Apache struct {
		// 主体配置
		Pulsar config `yaml:"pulsar" json:"pulsar" xml:"pulsar" toml:"pulsar"`
	} `yaml:"apache" json:"apache" xml:"apache" toml:"apache"`
}
