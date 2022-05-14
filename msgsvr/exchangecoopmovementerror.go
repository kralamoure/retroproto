// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCoopMovementError struct{}

func NewExchangeCoopMovementError(extra string) (ExchangeCoopMovementError, error) {
	var m ExchangeCoopMovementError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCoopMovementError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCoopMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCoopMovementError
}

func (m ExchangeCoopMovementError) MessageName() string {
	return "ExchangeCoopMovementError"
}

func (m ExchangeCoopMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCoopMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
