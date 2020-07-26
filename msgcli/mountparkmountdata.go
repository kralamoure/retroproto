// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountParkMountData struct{}

func (m MountParkMountData) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountParkMountData
}

func (m MountParkMountData) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountParkMountData) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
