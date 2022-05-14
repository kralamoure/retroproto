// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountMountParkBuy struct{}

func NewMountMountParkBuy(extra string) (MountMountParkBuy, error) {
	var m MountMountParkBuy

	if err := m.Deserialize(extra); err != nil {
		return MountMountParkBuy{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
