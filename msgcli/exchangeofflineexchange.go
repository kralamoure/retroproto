// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeOfflineExchange struct{}

func (m ExchangeOfflineExchange) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeOfflineExchange
}

func (m ExchangeOfflineExchange) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeOfflineExchange) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
