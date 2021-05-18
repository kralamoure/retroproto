package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountFree struct{}

func (m MountFree) ProtocolId() retroproto.MsgCliId {
	return retroproto.MountFree
}

func (m MountFree) Serialized() (string, error) {
	return "", nil
}

func (m *MountFree) Deserialize(extra string) error {
	return nil
}
