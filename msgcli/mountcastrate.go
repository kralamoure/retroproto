// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountCastrate struct{}

func (m MountCastrate) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountCastrate
}

func (m MountCastrate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountCastrate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
