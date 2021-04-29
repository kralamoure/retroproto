// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeGetCrafterForJob struct{}

func (m ExchangeGetCrafterForJob) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeGetCrafterForJob
}

func (m ExchangeGetCrafterForJob) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeGetCrafterForJob) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
