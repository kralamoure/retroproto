// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePlayerShopMovementError struct{}

func (m ExchangePlayerShopMovementError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangePlayerShopMovementError
}

func (m ExchangePlayerShopMovementError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangePlayerShopMovementError) Deserialize(extra string) error {
	return nil
}
