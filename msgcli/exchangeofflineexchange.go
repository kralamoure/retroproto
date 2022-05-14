// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeOfflineExchange struct{}

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
