// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type HousesKick struct{}

func (m HousesKick) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.HousesKick
}

func (m HousesKick) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesKick) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
