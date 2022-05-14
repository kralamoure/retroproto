// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type StoragesListRemove struct{}

func NewStoragesListRemove(extra string) (StoragesListRemove, error) {
	var m StoragesListRemove

	if err := m.Deserialize(extra); err != nil {
		return StoragesListRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m StoragesListRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.StoragesListRemove
}

func (m StoragesListRemove) MessageName() string {
	return "StoragesListRemove"
}

func (m StoragesListRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *StoragesListRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
