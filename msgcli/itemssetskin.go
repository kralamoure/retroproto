// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsSetSkin struct{}

func NewItemsSetSkin(extra string) (ItemsSetSkin, error) {
	var m ItemsSetSkin

	if err := m.Deserialize(extra); err != nil {
		return ItemsSetSkin{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsSetSkin) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsSetSkin
}

func (m ItemsSetSkin) MessageName() string {
	return "ItemsSetSkin"
}

func (m ItemsSetSkin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsSetSkin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
