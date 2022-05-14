// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountRemoveObjectInPark struct{}

func NewMountRemoveObjectInPark(extra string) (MountRemoveObjectInPark, error) {
	var m MountRemoveObjectInPark

	if err := m.Deserialize(extra); err != nil {
		return MountRemoveObjectInPark{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountRemoveObjectInPark) MessageId() retroproto.MsgCliId {
	return retroproto.MountRemoveObjectInPark
}

func (m MountRemoveObjectInPark) MessageName() string {
	return "MountRemoveObjectInPark"
}

func (m MountRemoveObjectInPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountRemoveObjectInPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
