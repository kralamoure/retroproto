// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeShop struct{}

func NewExchangeShop(extra string) (ExchangeShop, error) {
	var m ExchangeShop

	if err := m.Deserialize(extra); err != nil {
		return ExchangeShop{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeShop) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeShop
}

func (m ExchangeShop) MessageName() string {
	return "ExchangeShop"
}

func (m ExchangeShop) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeShop) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
