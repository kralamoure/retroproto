// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeMovementBuy struct{}

func (m ExchangeMovementBuy) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementBuy
}

func (m ExchangeMovementBuy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementBuy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
