package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountUnequip struct{}

func (m MountUnequip) ProtocolId() retroproto.MsgSvrId {
	return retroproto.MountUnequip
}

func (m MountUnequip) Serialized() (string, error) {
	return "", nil
}

func (m *MountUnequip) Deserialize(extra string) error {
	return nil
}
