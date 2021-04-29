package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountFree struct{}

func (m MountFree) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountFree
}

func (m MountFree) Serialized() (string, error) {
	return "", nil
}

func (m *MountFree) Deserialize(extra string) error {
	return nil
}
