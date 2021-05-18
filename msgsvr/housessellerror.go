// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesSellError struct{}

func (m HousesSellError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesSellError
}

func (m HousesSellError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesSellError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
