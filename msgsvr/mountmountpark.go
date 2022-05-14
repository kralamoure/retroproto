// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountMountPark struct{}

func (m MountMountPark) MessageId() retroproto.MsgSvrId {
	return retroproto.MountMountPark
}

func (m MountMountPark) MessageName() string {
	return "MountMountPark"
}

func (m MountMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
