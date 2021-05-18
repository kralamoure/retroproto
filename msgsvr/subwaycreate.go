// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SubwayCreate struct{}

func (m SubwayCreate) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SubwayCreate
}

func (m SubwayCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
