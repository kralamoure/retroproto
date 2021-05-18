// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLeave struct{}

func (m ExchangeLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeLeave
}

func (m ExchangeLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
