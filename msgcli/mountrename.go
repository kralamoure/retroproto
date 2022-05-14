package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountRename struct {
	Name string
}

func NewMountRename(extra string) (MountRename, error) {
	var m MountRename

	if err := m.Deserialize(extra); err != nil {
		return MountRename{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountRename) MessageId() retroproto.MsgCliId {
	return retroproto.MountRename
}

func (m MountRename) MessageName() string {
	return "MountRename"
}

func (m MountRename) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountRename) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
