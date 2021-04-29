// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangePlayerShopMovementSuccess struct{}

func (m ExchangePlayerShopMovementSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangePlayerShopMovementSuccess
}

func (m ExchangePlayerShopMovementSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangePlayerShopMovementSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
