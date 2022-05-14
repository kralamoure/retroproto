package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountEquipError struct {
	Reason rune
}

func NewMountEquipError(extra string) (MountEquipError, error) {
	var m MountEquipError

	if err := m.Deserialize(extra); err != nil {
		return MountEquipError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountEquipError) MessageId() retroproto.MsgSvrId {
	return retroproto.MountEquipError
}

func (m MountEquipError) MessageName() string {
	return "MountEquipError"
}

func (m MountEquipError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *MountEquipError) Deserialize(extra string) error {
	for _, reason := range extra {
		m.Reason = reason
		break
	}

	return nil
}
