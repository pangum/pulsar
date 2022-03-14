package pulsar

import (
	`strings`

	`github.com/pangum/pangu`
)

func newPulsar(config *pangu.Config) (client *Client, err error) {
	_panguConfig := new(panguConfig)
	if err = config.Load(_panguConfig); nil != err {
		return
	}

	pulsar := _panguConfig.Apache.Pulsar
	count := len(pulsar.Broker) + 1
	urls := make(map[string]string, count)
	_options := make(map[string]pulsarOptions, count)

	// 组装默认配置
	if `` != strings.TrimSpace(pulsar.URL) {
		urls[defaultLabel] = pulsar.URL
		_options[defaultLabel] = pulsar.Options
	}

	// 组装集群列表
	for _, _broker := range pulsar.Broker {
		urls[_broker.Label] = _broker.URL
		_options[_broker.Label] = _broker.Options
	}

	// 生成客户端封装实例
	client = newClient(urls, _options)

	return
}
