// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountParkMountData struct{}

func NewMountParkMountData(extra string) (MountParkMountData, error) {
	var m MountParkMountData

	if err := m.Deserialize(extra); err != nil {
		return MountParkMountData{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountParkMountData) MessageId() retroproto.MsgCliId {
	return retroproto.MountParkMountData
}

func (m MountParkMountData) MessageName() string {
	return "MountParkMountData"
}

func (m MountParkMountData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountParkMountData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
