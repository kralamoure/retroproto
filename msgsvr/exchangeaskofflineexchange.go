// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeAskOfflineExchange struct{}

func NewExchangeAskOfflineExchange(extra string) (ExchangeAskOfflineExchange, error) {
	var m ExchangeAskOfflineExchange

	if err := m.Deserialize(extra); err != nil {
		return ExchangeAskOfflineExchange{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeAskOfflineExchange) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeAskOfflineExchange
}

func (m ExchangeAskOfflineExchange) MessageName() string {
	return "ExchangeAskOfflineExchange"
}

func (m ExchangeAskOfflineExchange) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeAskOfflineExchange) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
