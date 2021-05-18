// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreTypeItemsMovementAdd struct{}

func (m ExchangeBigStoreTypeItemsMovementAdd) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreTypeItemsMovementAdd
}

func (m ExchangeBigStoreTypeItemsMovementAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeBigStoreTypeItemsMovementAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
