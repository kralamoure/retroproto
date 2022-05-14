package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeLeaveError struct{}

func NewExchangeLeaveError(extra string) (ExchangeLeaveError, error) {
	var m ExchangeLeaveError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeLeaveError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeLeaveError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLeaveError
}

func (m ExchangeLeaveError) MessageName() string {
	return "ExchangeLeaveError"
}

func (m ExchangeLeaveError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeLeaveError) Deserialize(extra string) error {
	return nil
}
