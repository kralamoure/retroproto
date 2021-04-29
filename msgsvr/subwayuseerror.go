// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayUseError struct{}

func (m SubwayUseError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SubwayUseError
}

func (m SubwayUseError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayUseError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
