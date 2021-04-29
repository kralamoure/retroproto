// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePlayerShopMovementSuccess struct{}

func (m ExchangePlayerShopMovementSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangePlayerShopMovementSuccess
}

func (m ExchangePlayerShopMovementSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangePlayerShopMovementSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
