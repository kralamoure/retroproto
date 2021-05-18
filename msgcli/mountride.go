package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type MountRide struct{}

func (m MountRide) ProtocolId() retroproto.MsgCliId {
	return retroproto.MountRide
}

func (m MountRide) Serialized() (string, error) {
	return "", nil
}

func (m *MountRide) Deserialize(extra string) error {
	return nil
}
