// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type SubwayRequestLeave struct{}

func (m SubwayRequestLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.SubwayRequestLeave
}

func (m SubwayRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
