// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountMountParkSell struct{}

func NewMountMountParkSell(extra string) (MountMountParkSell, error) {
	var m MountMountParkSell

	if err := m.Deserialize(extra); err != nil {
		return MountMountParkSell{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountMountParkSell) MessageId() retroproto.MsgCliId {
	return retroproto.MountMountParkSell
}

func (m MountMountParkSell) MessageName() string {
	return "MountMountParkSell"
}

func (m MountMountParkSell) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountMountParkSell) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
