package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountRide struct{}

func (m MountRide) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountRide
}

func (m MountRide) Serialized() (string, error) {
	return "", nil
}

func (m *MountRide) Deserialize(extra string) error {
	return nil
}
