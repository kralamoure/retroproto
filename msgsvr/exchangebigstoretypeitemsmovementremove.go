// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreTypeItemsMovementRemove struct{}

func NewExchangeBigStoreTypeItemsMovementRemove(extra string) (ExchangeBigStoreTypeItemsMovementRemove, error) {
	var m ExchangeBigStoreTypeItemsMovementRemove

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBigStoreTypeItemsMovementRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeBigStoreTypeItemsMovementRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreTypeItemsMovementRemove
}

func (m ExchangeBigStoreTypeItemsMovementRemove) MessageName() string {
	return "ExchangeBigStoreTypeItemsMovementRemove"
}

func (m ExchangeBigStoreTypeItemsMovementRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeBigStoreTypeItemsMovementRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
