// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreTypeItemsMovementRemove struct{}

func (m ExchangeBigStoreTypeItemsMovementRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreTypeItemsMovementRemove
}

func (m ExchangeBigStoreTypeItemsMovementRemove) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreTypeItemsMovementRemove) Deserialize(extra string) error {
	return nil
}
