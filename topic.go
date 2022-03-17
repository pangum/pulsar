package pulsar

import (
	`time`
)

type topic struct {
	// 主题
	Topic string `yaml:"topic" json:"topic" xml:"topic" toml:"topic"`
	// 主题列表
	Topics []string `yaml:"topics" json:"topics" xml:"topics" toml:"topics"`
	// 主题模式
	Pattern string `yaml:"pattern" json:"pattern" xml:"pattern" toml:"pattern"`
	// 属性
	Properties map[string]string `yaml:"properties" json:"properties" xml:"properties" toml:"properties"`
	// 超时
	Timeout time.Duration `json:"timeout" json:"timeout" xml:"timeout" toml:"timeout"`
}
