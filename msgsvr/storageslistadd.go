// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type StoragesListAdd struct{}

func NewStoragesListAdd(extra string) (StoragesListAdd, error) {
	var m StoragesListAdd

	if err := m.Deserialize(extra); err != nil {
		return StoragesListAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m StoragesListAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.StoragesListAdd
}

func (m StoragesListAdd) MessageName() string {
	return "StoragesListAdd"
}

func (m StoragesListAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *StoragesListAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
