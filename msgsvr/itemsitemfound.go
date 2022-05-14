// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsItemFound struct{}

func NewItemsItemFound(extra string) (ItemsItemFound, error) {
	var m ItemsItemFound

	if err := m.Deserialize(extra); err != nil {
		return ItemsItemFound{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsItemFound) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsItemFound
}

func (m ItemsItemFound) MessageName() string {
	return "ItemsItemFound"
}

func (m ItemsItemFound) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsItemFound) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
