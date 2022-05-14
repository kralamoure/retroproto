// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeMovementBuy struct{}

func NewExchangeMovementBuy(extra string) (ExchangeMovementBuy, error) {
	var m ExchangeMovementBuy

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMovementBuy{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMovementBuy) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementBuy
}

func (m ExchangeMovementBuy) MessageName() string {
	return "ExchangeMovementBuy"
}

func (m ExchangeMovementBuy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementBuy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
