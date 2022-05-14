// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreTypeItemsMovementRemove struct{}

func (m ExchangeBigStoreTypeItemsMovementRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreTypeItemsMovementRemove
}

func (m ExchangeBigStoreTypeItemsMovementRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeBigStoreTypeItemsMovementRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
