// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsChange struct{}

func NewItemsChange(extra string) (ItemsChange, error) {
	var m ItemsChange

	if err := m.Deserialize(extra); err != nil {
		return ItemsChange{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsChange) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsChange
}

func (m ItemsChange) MessageName() string {
	return "ItemsChange"
}

func (m ItemsChange) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsChange) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
