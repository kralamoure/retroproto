// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountKill struct{}

func (m MountKill) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountKill
}

func (m MountKill) Serialized() (string, error) {
	return "", nil
}

func (m *MountKill) Deserialize(extra string) error {
	return nil
}
