// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeKillMount struct{}

func (m ExchangeKillMount) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeKillMount
}

func (m ExchangeKillMount) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeKillMount) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
