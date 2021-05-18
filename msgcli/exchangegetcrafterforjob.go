// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeGetCrafterForJob struct{}

func (m ExchangeGetCrafterForJob) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeGetCrafterForJob
}

func (m ExchangeGetCrafterForJob) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeGetCrafterForJob) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
