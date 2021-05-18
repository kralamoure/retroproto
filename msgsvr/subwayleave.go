// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SubwayLeave struct{}

func (m SubwayLeave) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SubwayLeave
}

func (m SubwayLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
