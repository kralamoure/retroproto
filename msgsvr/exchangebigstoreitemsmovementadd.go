// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreItemsMovementAdd struct{}

func (m ExchangeBigStoreItemsMovementAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreItemsMovementAdd
}

func (m ExchangeBigStoreItemsMovementAdd) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreItemsMovementAdd) Deserialize(extra string) error {
	return nil
}
