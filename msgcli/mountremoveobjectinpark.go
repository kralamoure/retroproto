// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountRemoveObjectInPark struct{}

func (m MountRemoveObjectInPark) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountRemoveObjectInPark
}

func (m MountRemoveObjectInPark) Serialized() (string, error) {
	return "", nil
}

func (m *MountRemoveObjectInPark) Deserialize(extra string) error {
	return nil
}
