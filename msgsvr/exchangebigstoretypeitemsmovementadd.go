// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreTypeItemsMovementAdd struct{}

func (m ExchangeBigStoreTypeItemsMovementAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeBigStoreTypeItemsMovementAdd
}

func (m ExchangeBigStoreTypeItemsMovementAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeBigStoreTypeItemsMovementAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
