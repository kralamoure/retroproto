// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreItemsMovementAdd struct{}

func (m ExchangeBigStoreItemsMovementAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeBigStoreItemsMovementAdd
}

func (m ExchangeBigStoreItemsMovementAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeBigStoreItemsMovementAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
