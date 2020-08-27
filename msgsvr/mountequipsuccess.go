package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountEquipSuccess struct {
	MountData MountData
}

func (m MountEquipSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipSuccess
}

func (m MountEquipSuccess) Serialized() (string, error) {
	return m.MountData.Serialized()
}

func (m *MountEquipSuccess) Deserialize(extra string) error {
	return m.MountData.Deserialize(extra)
}
