// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangePlayerShopMovementError struct{}

func (m ExchangePlayerShopMovementError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangePlayerShopMovementError
}

func (m ExchangePlayerShopMovementError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangePlayerShopMovementError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
