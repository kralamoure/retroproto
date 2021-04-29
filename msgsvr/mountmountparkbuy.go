// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountMountParkBuy struct{}

func (m MountMountParkBuy) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountMountParkBuy
}

func (m MountMountParkBuy) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountMountParkBuy) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
