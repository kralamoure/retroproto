// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCoopMovementSuccess struct{}

func NewExchangeCoopMovementSuccess(extra string) (ExchangeCoopMovementSuccess, error) {
	var m ExchangeCoopMovementSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCoopMovementSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCoopMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCoopMovementSuccess
}

func (m ExchangeCoopMovementSuccess) MessageName() string {
	return "ExchangeCoopMovementSuccess"
}

func (m ExchangeCoopMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCoopMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
