// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeShop struct{}

func (m ExchangeShop) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeShop
}

func (m ExchangeShop) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeShop) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
