// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeLocalDistantError struct{}

func NewExchangeLocalDistantError(extra string) (ExchangeLocalDistantError, error) {
	var m ExchangeLocalDistantError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeLocalDistantError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeLocalDistantError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalDistantError
}

func (m ExchangeLocalDistantError) MessageName() string {
	return "ExchangeLocalDistantError"
}

func (m ExchangeLocalDistantError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalDistantError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
