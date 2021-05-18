package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountRename struct {
	Name string
}

func (m MountRename) ProtocolId() retroproto.MsgCliId {
	return retroproto.MountRename
}

func (m MountRename) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountRename) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
