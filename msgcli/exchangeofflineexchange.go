// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeOfflineExchange struct{}

func NewExchangeOfflineExchange(extra string) (ExchangeOfflineExchange, error) {
	var m ExchangeOfflineExchange

	if err := m.Deserialize(extra); err != nil {
		return ExchangeOfflineExchange{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeOfflineExchange) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeOfflineExchange
}

func (m ExchangeOfflineExchange) MessageName() string {
	return "ExchangeOfflineExchange"
}

func (m ExchangeOfflineExchange) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeOfflineExchange) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
