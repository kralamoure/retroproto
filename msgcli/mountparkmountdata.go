// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountParkMountData struct{}

func (m MountParkMountData) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountParkMountData
}

func (m MountParkMountData) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountParkMountData) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
