// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsItemUseCondition struct{}

func NewItemsItemUseCondition(extra string) (ItemsItemUseCondition, error) {
	var m ItemsItemUseCondition

	if err := m.Deserialize(extra); err != nil {
		return ItemsItemUseCondition{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsItemUseCondition) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsItemUseCondition
}

func (m ItemsItemUseCondition) MessageName() string {
	return "ItemsItemUseCondition"
}

func (m ItemsItemUseCondition) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsItemUseCondition) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
