package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountName struct {
	Name string
}

func NewMountName(extra string) (MountName, error) {
	var m MountName

	if err := m.Deserialize(extra); err != nil {
		return MountName{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountName) MessageId() retroproto.MsgSvrId {
	return retroproto.MountName
}

func (m MountName) MessageName() string {
	return "MountName"
}

func (m MountName) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountName) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
