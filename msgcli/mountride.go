package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountRide struct{}

func NewMountRide(extra string) (MountRide, error) {
	var m MountRide

	if err := m.Deserialize(extra); err != nil {
		return MountRide{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountRide) MessageId() retroproto.MsgCliId {
	return retroproto.MountRide
}

func (m MountRide) MessageName() string {
	return "MountRide"
}

func (m MountRide) Serialized() (string, error) {
	return "", nil
}

func (m *MountRide) Deserialize(extra string) error {
	return nil
}
