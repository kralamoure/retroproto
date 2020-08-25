package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountEquipAdd struct {
	MountData
}

func (m MountEquipAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipAdd
}

func (m MountEquipAdd) Serialized() (string, error) {
	return m.MountData.Serialized()
}

func (m *MountEquipAdd) Deserialize(extra string) error {
	return m.MountData.Deserialize(extra)
}
