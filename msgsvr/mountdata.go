package msgsvr

import (
	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/typ"
)

type MountData struct {
	Data typ.CommonMountData
}

func (m MountData) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountData
}

func (m MountData) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountData) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
