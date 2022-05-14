package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountFree struct{}

func NewMountFree(extra string) (MountFree, error) {
	var m MountFree

	if err := m.Deserialize(extra); err != nil {
		return MountFree{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountFree) MessageId() retroproto.MsgCliId {
	return retroproto.MountFree
}

func (m MountFree) MessageName() string {
	return "MountFree"
}

func (m MountFree) Serialized() (string, error) {
	return "", nil
}

func (m *MountFree) Deserialize(extra string) error {
	return nil
}
