// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreTypeItemsMovementAdd struct{}

func (m ExchangeBigStoreTypeItemsMovementAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreTypeItemsMovementAdd
}

func (m ExchangeBigStoreTypeItemsMovementAdd) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreTypeItemsMovementAdd) Deserialize(extra string) error {
	return nil
}
