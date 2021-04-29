// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type HousesBuy struct{}

func (m HousesBuy) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.HousesBuy
}

func (m HousesBuy) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesBuy) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
