// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SubwayUseError struct{}

func (m SubwayUseError) MessageId() retroproto.MsgSvrId {
	return retroproto.SubwayUseError
}

func (m SubwayUseError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayUseError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
