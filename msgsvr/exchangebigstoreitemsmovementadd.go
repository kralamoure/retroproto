// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreItemsMovementAdd struct{}

func (m ExchangeBigStoreItemsMovementAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreItemsMovementAdd
}

func (m ExchangeBigStoreItemsMovementAdd) MessageName() string {
	return "ExchangeBigStoreItemsMovementAdd"
}

func (m ExchangeBigStoreItemsMovementAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeBigStoreItemsMovementAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
