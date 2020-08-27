package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountUnequip struct{}

func (m MountUnequip) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountUnequip
}

func (m MountUnequip) Serialized() (string, error) {
	return "", nil
}

func (m *MountUnequip) Deserialize(extra string) error {
	return nil
}
