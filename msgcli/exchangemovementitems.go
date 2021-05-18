// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeMovementItems struct{}

func (m ExchangeMovementItems) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementItems
}

func (m ExchangeMovementItems) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementItems) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
