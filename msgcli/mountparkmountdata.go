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
	return "", nil
}

func (m *MountParkMountData) Deserialize(extra string) error {
	return nil
}
