package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountCastrate struct{}

func (m MountCastrate) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountCastrate
}

func (m MountCastrate) Serialized() (string, error) {
	return "", nil
}

func (m *MountCastrate) Deserialize(extra string) error {
	return nil
}
