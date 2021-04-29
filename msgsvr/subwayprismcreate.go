// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SubwayPrismCreate struct{}

func (m SubwayPrismCreate) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SubwayPrismCreate
}

func (m SubwayPrismCreate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SubwayPrismCreate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
