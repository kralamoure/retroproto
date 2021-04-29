// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountMountParkSell struct{}

func (m MountMountParkSell) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountMountParkSell
}

func (m MountMountParkSell) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountMountParkSell) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
