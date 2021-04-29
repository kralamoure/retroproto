// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type MountLeave struct{}

func (m MountLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountLeave
}

func (m MountLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
