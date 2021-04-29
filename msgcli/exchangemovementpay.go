// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeMovementPay struct{}

func (m ExchangeMovementPay) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeMovementPay
}

func (m ExchangeMovementPay) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeMovementPay) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
