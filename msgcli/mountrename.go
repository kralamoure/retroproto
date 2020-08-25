package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountRename struct {
	Name string
}

func (m MountRename) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountRename
}

func (m MountRename) Serialized() (string, error) {
	return m.Name, nil
}

func (m *MountRename) Deserialize(extra string) error {
	m.Name = extra
	return nil
}
