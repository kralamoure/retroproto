package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountCastrate struct{}

func NewMountCastrate(extra string) (MountCastrate, error) {
	var m MountCastrate

	if err := m.Deserialize(extra); err != nil {
		return MountCastrate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountCastrate) MessageId() retroproto.MsgCliId {
	return retroproto.MountCastrate
}

func (m MountCastrate) MessageName() string {
	return "MountCastrate"
}

func (m MountCastrate) Serialized() (string, error) {
	return "", nil
}

func (m *MountCastrate) Deserialize(extra string) error {
	return nil
}
