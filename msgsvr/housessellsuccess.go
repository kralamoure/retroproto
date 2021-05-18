// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesSellSuccess struct{}

func (m HousesSellSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesSellSuccess
}

func (m HousesSellSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesSellSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
