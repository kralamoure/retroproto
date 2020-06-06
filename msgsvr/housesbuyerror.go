// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesBuyError struct{}

func (m HousesBuyError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesBuyError
}

func (m HousesBuyError) Serialized() (string, error) {
	return "", nil
}

func (m *HousesBuyError) Deserialize(extra string) error {
	return nil
}
