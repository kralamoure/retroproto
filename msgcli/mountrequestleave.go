package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountRequestLeave struct{}

func (m MountRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountRequestLeave
}

func (m MountRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *MountRequestLeave) Deserialize(extra string) error {
	return nil
}
