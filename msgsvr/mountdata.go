package msgsvr

import (
	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type MountData struct {
	Data typ.CommonMountData
}

func (m MountData) ProtocolId() retroproto.MsgSvrId {
	return retroproto.MountData
}

func (m MountData) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountData) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
