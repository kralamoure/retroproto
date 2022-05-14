// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangePayMovementError struct{}

func NewExchangePayMovementError(extra string) (ExchangePayMovementError, error) {
	var m ExchangePayMovementError

	if err := m.Deserialize(extra); err != nil {
		return ExchangePayMovementError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePayMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangePayMovementError
}

func (m ExchangePayMovementError) MessageName() string {
	return "ExchangePayMovementError"
}

func (m ExchangePayMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePayMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
