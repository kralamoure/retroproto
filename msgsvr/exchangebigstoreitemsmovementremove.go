// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreItemsMovementRemove struct{}

func (m ExchangeBigStoreItemsMovementRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreItemsMovementRemove
}

func (m ExchangeBigStoreItemsMovementRemove) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeBigStoreItemsMovementRemove) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
