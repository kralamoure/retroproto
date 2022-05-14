// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsDropError struct{}

func NewItemsDropError(extra string) (ItemsDropError, error) {
	var m ItemsDropError

	if err := m.Deserialize(extra); err != nil {
		return ItemsDropError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsDropError) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsDropError
}

func (m ItemsDropError) MessageName() string {
	return "ItemsDropError"
}

func (m ItemsDropError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsDropError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
