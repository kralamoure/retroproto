// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountMountPark struct{}

func (m MountMountPark) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountMountPark
}

func (m MountMountPark) Serialized() (string, error) {
	return "", nil
}

func (m *MountMountPark) Deserialize(extra string) error {
	return nil
}
