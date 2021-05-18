// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type WaypointsLeave struct{}

func (m WaypointsLeave) ProtocolId() retroproto.MsgSvrId {
	return retroproto.WaypointsLeave
}

func (m WaypointsLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
