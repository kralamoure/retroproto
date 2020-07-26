// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountName struct{}

func (m MountName) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountName
}

func (m MountName) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountName) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
