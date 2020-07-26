// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SubwayUseError struct{}

func (m SubwayUseError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SubwayUseError
}

func (m SubwayUseError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SubwayUseError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
