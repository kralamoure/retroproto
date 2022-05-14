// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeMovementSell struct{}

func NewExchangeMovementSell(extra string) (ExchangeMovementSell, error) {
	var m ExchangeMovementSell

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMovementSell{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMovementSell) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementSell
}

func (m ExchangeMovementSell) MessageName() string {
	return "ExchangeMovementSell"
}

func (m ExchangeMovementSell) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementSell) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
