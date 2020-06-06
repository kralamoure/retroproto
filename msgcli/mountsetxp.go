// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountSetXP struct{}

func (m MountSetXP) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountSetXP
}

func (m MountSetXP) Serialized() (string, error) {
	return "", nil
}

func (m *MountSetXP) Deserialize(extra string) error {
	return nil
}
