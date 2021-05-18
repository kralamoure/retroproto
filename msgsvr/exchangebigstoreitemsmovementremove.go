// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreItemsMovementRemove struct{}

func (m ExchangeBigStoreItemsMovementRemove) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreItemsMovementRemove
}

func (m ExchangeBigStoreItemsMovementRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeBigStoreItemsMovementRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
