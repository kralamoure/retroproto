// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeStorageMovementError struct{}

func NewExchangeStorageMovementError(extra string) (ExchangeStorageMovementError, error) {
	var m ExchangeStorageMovementError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeStorageMovementError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeStorageMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeStorageMovementError
}

func (m ExchangeStorageMovementError) MessageName() string {
	return "ExchangeStorageMovementError"
}

func (m ExchangeStorageMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeStorageMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
