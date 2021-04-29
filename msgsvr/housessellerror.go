// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesSellError struct{}

func (m HousesSellError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesSellError
}

func (m HousesSellError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesSellError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
