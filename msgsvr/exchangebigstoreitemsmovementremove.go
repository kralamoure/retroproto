// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreItemsMovementRemove struct{}

func NewExchangeBigStoreItemsMovementRemove(extra string) (ExchangeBigStoreItemsMovementRemove, error) {
	var m ExchangeBigStoreItemsMovementRemove

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBigStoreItemsMovementRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeBigStoreItemsMovementRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreItemsMovementRemove
}

func (m ExchangeBigStoreItemsMovementRemove) MessageName() string {
	return "ExchangeBigStoreItemsMovementRemove"
}

func (m ExchangeBigStoreItemsMovementRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeBigStoreItemsMovementRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
