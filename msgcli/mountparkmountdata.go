// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountParkMountData struct{}

func (m MountParkMountData) ProtocolId() retroproto.MsgCliId {
	return retroproto.MountParkMountData
}

func (m MountParkMountData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountParkMountData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
