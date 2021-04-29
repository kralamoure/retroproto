package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type MountEquipError struct {
	Reason rune
}

func (m MountEquipError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountEquipError
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
