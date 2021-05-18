// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type HousesBuy struct{}

func (m HousesBuy) ProtocolId() retroproto.MsgCliId {
	return retroproto.HousesBuy
}

func (m HousesBuy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesBuy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
