// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsDissociate struct{}

func NewItemsDissociate(extra string) (ItemsDissociate, error) {
	var m ItemsDissociate

	if err := m.Deserialize(extra); err != nil {
		return ItemsDissociate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsDissociate) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsDissociate
}

func (m ItemsDissociate) MessageName() string {
	return "ItemsDissociate"
}

func (m ItemsDissociate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsDissociate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
