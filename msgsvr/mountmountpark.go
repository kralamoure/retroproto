// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountMountPark struct{}

func NewMountMountPark(extra string) (MountMountPark, error) {
	var m MountMountPark

	if err := m.Deserialize(extra); err != nil {
		return MountMountPark{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountMountPark) MessageId() retroproto.MsgSvrId {
	return retroproto.MountMountPark
}

func (m MountMountPark) MessageName() string {
	return "MountMountPark"
}

func (m MountMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
