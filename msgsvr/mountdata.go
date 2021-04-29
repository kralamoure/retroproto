package msgsvr

import (
	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/typ"
)

type MountData struct {
	Data typ.CommonMountData
}

func (m MountData) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountData
}

func (m MountData) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountData) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
