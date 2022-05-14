// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeMovementPay struct{}

func (m ExchangeMovementPay) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementPay
}

func (m ExchangeMovementPay) MessageName() string {
	return "ExchangeMovementPay"
}

func (m ExchangeMovementPay) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementPay) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
