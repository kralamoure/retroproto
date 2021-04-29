// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountRequestMountParkBuy struct{}

func (m MountRequestMountParkBuy) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountRequestMountParkBuy
}

func (m MountRequestMountParkBuy) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountRequestMountParkBuy) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
