// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangePayMovementSuccess struct{}

func NewExchangePayMovementSuccess(extra string) (ExchangePayMovementSuccess, error) {
	var m ExchangePayMovementSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangePayMovementSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePayMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangePayMovementSuccess
}

func (m ExchangePayMovementSuccess) MessageName() string {
	return "ExchangePayMovementSuccess"
}

func (m ExchangePayMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePayMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
