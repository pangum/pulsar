package pulsar

import (
	`time`
)

type timeout struct {
	// 操作
	Operation time.Duration `default:"30s" yaml:"operation" json:"operation" xml:"operation" toml:"operation"`
	// 连接
	Connection time.Duration `default:"30s" xml:"connection" json:"connection" xml:"connection" toml:"connection"`
}
