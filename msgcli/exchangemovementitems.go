// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeMovementItems struct{}

func NewExchangeMovementItems(extra string) (ExchangeMovementItems, error) {
	var m ExchangeMovementItems

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMovementItems{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMovementItems) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementItems
}

func (m ExchangeMovementItems) MessageName() string {
	return "ExchangeMovementItems"
}

func (m ExchangeMovementItems) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementItems) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
