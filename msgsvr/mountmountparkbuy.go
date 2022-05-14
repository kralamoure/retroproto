// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountMountParkBuy struct{}

func (m MountMountParkBuy) MessageId() retroproto.MsgSvrId {
	return retroproto.MountMountParkBuy
}

func (m MountMountParkBuy) MessageName() string {
	return "MountMountParkBuy"
}

func (m MountMountParkBuy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountMountParkBuy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
