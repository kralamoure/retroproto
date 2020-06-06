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
	return "", nil
}

func (m *HousesBuy) Deserialize(extra string) error {
	return nil
}
