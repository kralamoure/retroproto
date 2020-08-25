package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountCastrate struct{}

func (m MountCastrate) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountCastrate
}

func (m MountCastrate) Serialized() (string, error) {
	return "", nil
}

func (m *MountCastrate) Deserialize(extra string) error {
	return nil
}
