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
	return "", nil
}

func (m *ExchangePlayerShopMovementSuccess) Deserialize(extra string) error {
	return nil
}
