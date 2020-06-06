// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesSellSuccess struct{}

func (m HousesSellSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesSellSuccess
}

func (m HousesSellSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *HousesSellSuccess) Deserialize(extra string) error {
	return nil
}
