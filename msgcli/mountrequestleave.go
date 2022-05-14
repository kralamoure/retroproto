package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountRequestLeave struct{}

func NewMountRequestLeave(extra string) (MountRequestLeave, error) {
	var m MountRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return MountRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.MountRequestLeave
}

func (m MountRequestLeave) MessageName() string {
	return "MountRequestLeave"
}

func (m MountRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *MountRequestLeave) Deserialize(extra string) error {
	return nil
}
