// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountRequestMountParkBuy struct{}

func (m MountRequestMountParkBuy) ProtocolId() retroproto.MsgCliId {
	return retroproto.MountRequestMountParkBuy
}

func (m MountRequestMountParkBuy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountRequestMountParkBuy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
