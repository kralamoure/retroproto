// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayCreate struct{}

func (m SubwayCreate) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SubwayCreate
}

func (m SubwayCreate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayCreate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
