// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountRemoveObjectInPark struct{}

func (m MountRemoveObjectInPark) MessageId() retroproto.MsgCliId {
	return retroproto.MountRemoveObjectInPark
}

func (m MountRemoveObjectInPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *MountRemoveObjectInPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
