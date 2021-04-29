// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayPrismCreate struct{}

func (m SubwayPrismCreate) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SubwayPrismCreate
}

func (m SubwayPrismCreate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayPrismCreate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
