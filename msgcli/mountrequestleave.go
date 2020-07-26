// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountRequestLeave struct{}

func (m MountRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountRequestLeave
}

func (m MountRequestLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountRequestLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
