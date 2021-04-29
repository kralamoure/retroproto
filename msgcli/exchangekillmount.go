// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeKillMount struct{}

func (m ExchangeKillMount) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeKillMount
}

func (m ExchangeKillMount) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeKillMount) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
