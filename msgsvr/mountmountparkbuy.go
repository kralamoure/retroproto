// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type MountMountParkBuy struct{}

func (m MountMountParkBuy) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountMountParkBuy
}

func (m MountMountParkBuy) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountMountParkBuy) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
