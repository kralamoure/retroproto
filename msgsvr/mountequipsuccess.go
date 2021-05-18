package msgsvr

import (
	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type MountEquipSuccess struct {
	Data typ.CommonMountData
}

func (m MountEquipSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.MountEquipSuccess
}

func (m MountEquipSuccess) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountEquipSuccess) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
