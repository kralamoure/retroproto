package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountEquipError struct {
	Reason rune
}

func (m MountEquipError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipError
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
