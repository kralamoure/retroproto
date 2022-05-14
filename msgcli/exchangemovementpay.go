// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeMovementPay struct{}

func NewExchangeMovementPay(extra string) (ExchangeMovementPay, error) {
	var m ExchangeMovementPay

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMovementPay{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
