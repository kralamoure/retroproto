// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreTypeItemsMovementAdd struct{}

func NewExchangeBigStoreTypeItemsMovementAdd(extra string) (ExchangeBigStoreTypeItemsMovementAdd, error) {
	var m ExchangeBigStoreTypeItemsMovementAdd

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBigStoreTypeItemsMovementAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeBigStoreTypeItemsMovementAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreTypeItemsMovementAdd
}

func (m ExchangeBigStoreTypeItemsMovementAdd) MessageName() string {
	return "ExchangeBigStoreTypeItemsMovementAdd"
}

func (m ExchangeBigStoreTypeItemsMovementAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeBigStoreTypeItemsMovementAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
