package pulsar

import `github.com/pangum/pangu`

func init() {
	pangu.New().Musts(newPulsar)
}
