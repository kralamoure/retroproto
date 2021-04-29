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
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeBigStoreItemsMovementAdd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
