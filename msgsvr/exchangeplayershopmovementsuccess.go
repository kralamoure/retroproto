// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangePlayerShopMovementSuccess struct{}

func (m ExchangePlayerShopMovementSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangePlayerShopMovementSuccess
}

func (m ExchangePlayerShopMovementSuccess) MessageName() string {
	return "ExchangePlayerShopMovementSuccess"
}

func (m ExchangePlayerShopMovementSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePlayerShopMovementSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
