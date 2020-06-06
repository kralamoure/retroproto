// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type HousesSell struct{}

func (m HousesSell) ProtocolId() d1proto.MsgCliId {
	return d1proto.HousesSell
}

func (m HousesSell) Serialized() (string, error) {
	return "", nil
}

func (m *HousesSell) Deserialize(extra string) error {
	return nil
}
