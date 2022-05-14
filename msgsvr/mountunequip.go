package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountUnequip struct{}

func NewMountUnequip(extra string) (MountUnequip, error) {
	var m MountUnequip

	if err := m.Deserialize(extra); err != nil {
		return MountUnequip{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountUnequip) MessageId() retroproto.MsgSvrId {
	return retroproto.MountUnequip
}

func (m MountUnequip) MessageName() string {
	return "MountUnequip"
}

func (m MountUnequip) Serialized() (string, error) {
	return "", nil
}

func (m *MountUnequip) Deserialize(extra string) error {
	return nil
}
