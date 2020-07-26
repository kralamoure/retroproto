// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountRename struct{}

func (m MountRename) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountRename
}

func (m MountRename) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountRename) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
