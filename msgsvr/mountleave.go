// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountLeave struct{}

func (m MountLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountLeave
}

func (m MountLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
