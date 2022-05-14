// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeAskOfflineExchange struct{}

func (m ExchangeAskOfflineExchange) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeAskOfflineExchange
}

func (m ExchangeAskOfflineExchange) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeAskOfflineExchange) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
