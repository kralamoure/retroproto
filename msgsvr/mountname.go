package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type MountName struct {
	Name string
}

func (m MountName) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountName
}

func (m MountName) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountName) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
