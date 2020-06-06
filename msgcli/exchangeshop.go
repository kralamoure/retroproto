// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeShop struct{}

func (m ExchangeShop) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeShop
}

func (m ExchangeShop) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeShop) Deserialize(extra string) error {
	return nil
}
