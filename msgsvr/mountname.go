package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountName struct {
	Name string
}

func (m MountName) MessageId() retroproto.MsgSvrId {
	return retroproto.MountName
}

func (m MountName) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountName) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
