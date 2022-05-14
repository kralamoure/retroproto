// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeRequestAskOfflineExchange struct{}

func NewExchangeRequestAskOfflineExchange(extra string) (ExchangeRequestAskOfflineExchange, error) {
	var m ExchangeRequestAskOfflineExchange

	if err := m.Deserialize(extra); err != nil {
		return ExchangeRequestAskOfflineExchange{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeRequestAskOfflineExchange) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeRequestAskOfflineExchange
}

func (m ExchangeRequestAskOfflineExchange) MessageName() string {
	return "ExchangeRequestAskOfflineExchange"
}

func (m ExchangeRequestAskOfflineExchange) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRequestAskOfflineExchange) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
