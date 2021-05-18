// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountSetXP struct{}

func (m MountSetXP) ProtocolId() retroproto.MsgCliId {
	return retroproto.MountSetXP
}

func (m MountSetXP) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountSetXP) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
