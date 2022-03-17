package pulsar

import (
	`fmt`
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
	clients := make(map[string]topic, count)

	// 组装默认配置
	if `` != strings.TrimSpace(pulsar.URL) {
		urls[defaultLabel] = pulsar.URL
		_options[defaultLabel] = pulsar.Options
		parseClients(defaultLabel, clients, pulsar.Clients...)
	}

	// 组装集群列表
	for _, _broker := range pulsar.Broker {
		urls[_broker.Label] = _broker.URL
		_options[_broker.Label] = _broker.Options
		parseClients(_broker.Label, clients, _broker.Clients...)
	}

	// 生成客户端封装实例
	client = newClient(urls, _options, clients)

	return
}

func parseClients(label string, clients map[string]topic, configs ...topicConfig) {
	for _, conf := range configs {
		var key string
		if `` == conf.Key {
			key = fmt.Sprintf(keyFormatter, label, defaultKey)
		} else {
			key = fmt.Sprintf(keyFormatter, label, conf.Key)
		}
		clients[key] = conf.topic
	}
}
