// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayLeave struct{}

func (m SubwayLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SubwayLeave
}

func (m SubwayLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
