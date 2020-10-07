package msgsvr

import (
	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/typ"
)

type MountEquipSuccess struct {
	Data typ.CommonMountData
}

func (m MountEquipSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipSuccess
}

func (m MountEquipSuccess) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountEquipSuccess) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
