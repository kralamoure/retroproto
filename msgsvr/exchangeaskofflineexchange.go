// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeAskOfflineExchange struct{}

func (m ExchangeAskOfflineExchange) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeAskOfflineExchange
}

func (m ExchangeAskOfflineExchange) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeAskOfflineExchange) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
