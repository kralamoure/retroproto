package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type MountUnequip struct{}

func (m MountUnequip) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.MountUnequip
}

func (m MountUnequip) Serialized() (string, error) {
	return "", nil
}

func (m *MountUnequip) Deserialize(extra string) error {
	return nil
}
