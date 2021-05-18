// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type HousesKick struct{}

func (m HousesKick) ProtocolId() retroproto.MsgCliId {
	return retroproto.HousesKick
}

func (m HousesKick) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesKick) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
