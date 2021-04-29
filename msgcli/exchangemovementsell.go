// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeMovementSell struct{}

func (m ExchangeMovementSell) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeMovementSell
}

func (m ExchangeMovementSell) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeMovementSell) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
