// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountMountPark struct{}

func (m MountMountPark) ProtocolId() retroproto.MsgSvrId {
	return retroproto.MountMountPark
}

func (m MountMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
