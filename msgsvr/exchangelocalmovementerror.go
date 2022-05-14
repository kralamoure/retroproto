// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeLocalMovementError struct{}

func NewExchangeLocalMovementError(extra string) (ExchangeLocalMovementError, error) {
	var m ExchangeLocalMovementError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeLocalMovementError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeLocalMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalMovementError
}

func (m ExchangeLocalMovementError) MessageName() string {
	return "ExchangeLocalMovementError"
}

func (m ExchangeLocalMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
