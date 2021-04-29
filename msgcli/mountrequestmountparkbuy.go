// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountRequestMountParkBuy struct{}

func (m MountRequestMountParkBuy) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountRequestMountParkBuy
}

func (m MountRequestMountParkBuy) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountRequestMountParkBuy) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
