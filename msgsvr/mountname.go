package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountName struct {
	Name string
}

func (m MountName) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountName
}

func (m MountName) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountName) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
