// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type StoragesLockedProperty struct{}

func NewStoragesLockedProperty(extra string) (StoragesLockedProperty, error) {
	var m StoragesLockedProperty

	if err := m.Deserialize(extra); err != nil {
		return StoragesLockedProperty{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m StoragesLockedProperty) MessageId() retroproto.MsgSvrId {
	return retroproto.StoragesLockedProperty
}

func (m StoragesLockedProperty) MessageName() string {
	return "StoragesLockedProperty"
}

func (m StoragesLockedProperty) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *StoragesLockedProperty) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
