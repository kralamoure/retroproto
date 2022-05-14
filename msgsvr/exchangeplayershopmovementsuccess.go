// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangePlayerShopMovementSuccess struct{}

func NewExchangePlayerShopMovementSuccess(extra string) (ExchangePlayerShopMovementSuccess, error) {
	var m ExchangePlayerShopMovementSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangePlayerShopMovementSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangePlayerShopMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangePlayerShopMovementSuccess
}

func (m ExchangePlayerShopMovementSuccess) MessageName() string {
	return "ExchangePlayerShopMovementSuccess"
}

func (m ExchangePlayerShopMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePlayerShopMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
