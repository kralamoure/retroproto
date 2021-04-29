// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeAskOfflineExchange struct{}

func (m ExchangeAskOfflineExchange) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeAskOfflineExchange
}

func (m ExchangeAskOfflineExchange) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeAskOfflineExchange) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
