package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountFree struct{}

func (m MountFree) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountFree
}

func (m MountFree) Serialized() (string, error) {
	return "", nil
}

func (m *MountFree) Deserialize(extra string) error {
	return nil
}
