// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesBuySuccess struct{}

func (m HousesBuySuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesBuySuccess
}

func (m HousesBuySuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesBuySuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
