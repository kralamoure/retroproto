// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountLeave struct{}

func (m MountLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.MountLeave
}

func (m MountLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
