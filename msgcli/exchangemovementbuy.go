// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeMovementBuy struct{}

func (m ExchangeMovementBuy) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeMovementBuy
}

func (m ExchangeMovementBuy) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeMovementBuy) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
