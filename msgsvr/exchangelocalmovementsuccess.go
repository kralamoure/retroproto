// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeLocalMovementSuccess struct{}

func NewExchangeLocalMovementSuccess(extra string) (ExchangeLocalMovementSuccess, error) {
	var m ExchangeLocalMovementSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeLocalMovementSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeLocalMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalMovementSuccess
}

func (m ExchangeLocalMovementSuccess) MessageName() string {
	return "ExchangeLocalMovementSuccess"
}

func (m ExchangeLocalMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
