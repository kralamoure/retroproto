// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeRequestAskOfflineExchange struct{}

func (m ExchangeRequestAskOfflineExchange) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeRequestAskOfflineExchange
}

func (m ExchangeRequestAskOfflineExchange) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRequestAskOfflineExchange) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
