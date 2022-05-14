package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCreateError struct{}

func NewExchangeCreateError(extra string) (ExchangeCreateError, error) {
	var m ExchangeCreateError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCreateError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCreateError
}

func (m ExchangeCreateError) MessageName() string {
	return "ExchangeCreateError"
}

func (m ExchangeCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeCreateError) Deserialize(extra string) error {
	return nil
}
