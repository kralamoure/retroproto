// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeGetCrafterForJob struct{}

func (m ExchangeGetCrafterForJob) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeGetCrafterForJob
}

func (m ExchangeGetCrafterForJob) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeGetCrafterForJob) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
