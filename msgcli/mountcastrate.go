package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountCastrate struct{}

func (m MountCastrate) MessageId() retroproto.MsgCliId {
	return retroproto.MountCastrate
}

func (m MountCastrate) Serialized() (string, error) {
	return "", nil
}

func (m *MountCastrate) Deserialize(extra string) error {
	return nil
}
