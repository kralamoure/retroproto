// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangePlayerShopMovementError struct{}

func (m ExchangePlayerShopMovementError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangePlayerShopMovementError
}

func (m ExchangePlayerShopMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePlayerShopMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
