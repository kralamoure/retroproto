package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type MountRide struct{}

func (m MountRide) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.MountRide
}

func (m MountRide) Serialized() (string, error) {
	return "", nil
}

func (m *MountRide) Deserialize(extra string) error {
	return nil
}
