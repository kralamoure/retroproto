// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SubwayPrismLeave struct{}

func (m SubwayPrismLeave) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SubwayPrismLeave
}

func (m SubwayPrismLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayPrismLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
