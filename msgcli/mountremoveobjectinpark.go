// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountRemoveObjectInPark struct{}

func (m MountRemoveObjectInPark) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountRemoveObjectInPark
}

func (m MountRemoveObjectInPark) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *MountRemoveObjectInPark) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
