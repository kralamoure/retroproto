// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreItemsMovementAdd struct{}

func NewExchangeBigStoreItemsMovementAdd(extra string) (ExchangeBigStoreItemsMovementAdd, error) {
	var m ExchangeBigStoreItemsMovementAdd

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBigStoreItemsMovementAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
