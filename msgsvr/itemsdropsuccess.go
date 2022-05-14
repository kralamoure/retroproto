// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsDropSuccess struct{}

func NewItemsDropSuccess(extra string) (ItemsDropSuccess, error) {
	var m ItemsDropSuccess

	if err := m.Deserialize(extra); err != nil {
		return ItemsDropSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsDropSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsDropSuccess
}

func (m ItemsDropSuccess) MessageName() string {
	return "ItemsDropSuccess"
}

func (m ItemsDropSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsDropSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
