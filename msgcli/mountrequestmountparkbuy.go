// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountRequestMountParkBuy struct{}

func NewMountRequestMountParkBuy(extra string) (MountRequestMountParkBuy, error) {
	var m MountRequestMountParkBuy

	if err := m.Deserialize(extra); err != nil {
		return MountRequestMountParkBuy{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountRequestMountParkBuy) MessageId() retroproto.MsgCliId {
	return retroproto.MountRequestMountParkBuy
}

func (m MountRequestMountParkBuy) MessageName() string {
	return "MountRequestMountParkBuy"
}

func (m MountRequestMountParkBuy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountRequestMountParkBuy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
