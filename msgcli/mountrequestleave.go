package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountRequestLeave struct{}

func (m MountRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.MountRequestLeave
}

func (m MountRequestLeave) MessageName() string {
	return "MountRequestLeave"
}

func (m MountRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *MountRequestLeave) Deserialize(extra string) error {
	return nil
}
