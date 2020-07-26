// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type HousesKick struct{}

func (m HousesKick) ProtocolId() d1proto.MsgCliId {
	return d1proto.HousesKick
}

func (m HousesKick) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesKick) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
