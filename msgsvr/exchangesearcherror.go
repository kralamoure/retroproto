package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeSearchError struct{}

func NewExchangeSearchError(extra string) (ExchangeSearchError, error) {
	var m ExchangeSearchError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeSearchError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeSearchError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeSearchError
}

func (m ExchangeSearchError) MessageName() string {
	return "ExchangeSearchError"
}

func (m ExchangeSearchError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchError) Deserialize(extra string) error {
	return nil
}
