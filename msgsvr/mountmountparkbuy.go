// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountMountParkBuy struct{}

func (m MountMountParkBuy) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountMountParkBuy
}

func (m MountMountParkBuy) Serialized() (string, error) {
	return "", nil
}

func (m *MountMountParkBuy) Deserialize(extra string) error {
	return nil
}
