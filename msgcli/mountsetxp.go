// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountSetXP struct{}

func (m MountSetXP) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountSetXP
}

func (m MountSetXP) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountSetXP) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
