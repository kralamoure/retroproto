package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountEquipError struct{}

func (m MountEquipError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipError
}

func (m MountEquipError) Serialized() (string, error) {
	return "", nil
}

func (m *MountEquipError) Deserialize(extra string) error {
	return nil
}
