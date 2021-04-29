// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreItemsMovementRemove struct{}

func (m ExchangeBigStoreItemsMovementRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeBigStoreItemsMovementRemove
}

func (m ExchangeBigStoreItemsMovementRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeBigStoreItemsMovementRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
