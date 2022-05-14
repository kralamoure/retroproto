// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangePlayerShopMovementError struct{}

func NewExchangePlayerShopMovementError(extra string) (ExchangePlayerShopMovementError, error) {
	var m ExchangePlayerShopMovementError

	if err := m.Deserialize(extra); err != nil {
		return ExchangePlayerShopMovementError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePlayerShopMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangePlayerShopMovementError
}

func (m ExchangePlayerShopMovementError) MessageName() string {
	return "ExchangePlayerShopMovementError"
}

func (m ExchangePlayerShopMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePlayerShopMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
