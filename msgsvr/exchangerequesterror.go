package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/enum"
)

type ExchangeRequestError struct {
	Reason rune
}

func NewExchangeRequestError(extra string) (ExchangeRequestError, error) {
	var m ExchangeRequestError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeRequestError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeRequestError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeRequestError
}

func (m ExchangeRequestError) MessageName() string {
	return "ExchangeRequestError"
}

func (m ExchangeRequestError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ExchangeRequestError) Deserialize(extra string) error {
	if extra == "" {
		m.Reason = enum.ExchangeRequestErrorReason.Default
	} else {
		m.Reason = rune(extra[0])
	}
	return nil
}
