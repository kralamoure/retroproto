// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SubwayCreate struct{}

func (m SubwayCreate) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SubwayCreate
}

func (m SubwayCreate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SubwayCreate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
