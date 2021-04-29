// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeShop struct{}

func (m ExchangeShop) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeShop
}

func (m ExchangeShop) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeShop) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
