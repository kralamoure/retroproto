// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreTypeItemsMovementRemove struct{}

func (m ExchangeBigStoreTypeItemsMovementRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeBigStoreTypeItemsMovementRemove
}

func (m ExchangeBigStoreTypeItemsMovementRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeBigStoreTypeItemsMovementRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
