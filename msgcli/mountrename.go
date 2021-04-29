package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountRename struct {
	Name string
}

func (m MountRename) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountRename
}

func (m MountRename) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountRename) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
