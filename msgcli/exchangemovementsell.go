// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeMovementSell struct{}

func (m ExchangeMovementSell) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementSell
}

func (m ExchangeMovementSell) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementSell) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
