package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountEquipError struct {
	Reason rune
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
