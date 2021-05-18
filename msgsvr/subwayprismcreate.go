// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SubwayPrismCreate struct{}

func (m SubwayPrismCreate) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SubwayPrismCreate
}

func (m SubwayPrismCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayPrismCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
