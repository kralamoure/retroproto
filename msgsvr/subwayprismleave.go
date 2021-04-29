// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayPrismLeave struct{}

func (m SubwayPrismLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SubwayPrismLeave
}

func (m SubwayPrismLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayPrismLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
