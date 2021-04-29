// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeRequestAskOfflineExchange struct{}

func (m ExchangeRequestAskOfflineExchange) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeRequestAskOfflineExchange
}

func (m ExchangeRequestAskOfflineExchange) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeRequestAskOfflineExchange) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
