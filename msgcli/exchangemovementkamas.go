// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeMovementKamas struct{}

func NewExchangeMovementKamas(extra string) (ExchangeMovementKamas, error) {
	var m ExchangeMovementKamas

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMovementKamas{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMovementKamas) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeMovementKamas
}

func (m ExchangeMovementKamas) MessageName() string {
	return "ExchangeMovementKamas"
}

func (m ExchangeMovementKamas) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMovementKamas) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
