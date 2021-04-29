package msgsvr

import (
	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/typ"
)

type MountEquipSuccess struct {
	Data typ.CommonMountData
}

func (m MountEquipSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountEquipSuccess
}

func (m MountEquipSuccess) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountEquipSuccess) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
