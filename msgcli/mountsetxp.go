// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountSetXP struct{}

func NewMountSetXP(extra string) (MountSetXP, error) {
	var m MountSetXP

	if err := m.Deserialize(extra); err != nil {
		return MountSetXP{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountSetXP) MessageId() retroproto.MsgCliId {
	return retroproto.MountSetXP
}

func (m MountSetXP) MessageName() string {
	return "MountSetXP"
}

func (m MountSetXP) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountSetXP) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
