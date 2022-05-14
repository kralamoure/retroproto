package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountUnequip struct{}

func (m MountUnequip) MessageId() retroproto.MsgSvrId {
	return retroproto.MountUnequip
}

func (m MountUnequip) MessageName() string {
	return "MountUnequip"
}

func (m MountUnequip) Serialized() (string, error) {
	return "", nil
}

func (m *MountUnequip) Deserialize(extra string) error {
	return nil
}
