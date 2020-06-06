package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountUnEquipRemove struct{}

func (m MountUnEquipRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipRemove
}

func (m MountUnEquipRemove) Serialized() (string, error) {
	return "", nil
}

func (m *MountUnEquipRemove) Deserialize(extra string) error {
	return nil
}
