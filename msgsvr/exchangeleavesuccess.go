package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeLeaveSuccess struct {
	TypePlayerExchange bool
}

func NewExchangeLeaveSuccess(extra string) (ExchangeLeaveSuccess, error) {
	var m ExchangeLeaveSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeLeaveSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeLeaveSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLeaveSuccess
}

func (m ExchangeLeaveSuccess) MessageName() string {
	return "ExchangeLeaveSuccess"
}

func (m ExchangeLeaveSuccess) Serialized() (string, error) {
	if m.TypePlayerExchange {
		return "a", nil
	}
	return "", nil
}

func (m *ExchangeLeaveSuccess) Deserialize(extra string) error {
	if extra == "a" {
		m.TypePlayerExchange = true
	}
	return nil
}
