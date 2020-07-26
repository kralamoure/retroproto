// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type HousesBuy struct{}

func (m HousesBuy) ProtocolId() d1proto.MsgCliId {
	return d1proto.HousesBuy
}

func (m HousesBuy) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesBuy) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
