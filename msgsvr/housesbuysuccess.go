// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesBuySuccess struct{}

func (m HousesBuySuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesBuySuccess
}

func (m HousesBuySuccess) Serialized() (string, error) {
	return "", nil
}

func (m *HousesBuySuccess) Deserialize(extra string) error {
	return nil
}
