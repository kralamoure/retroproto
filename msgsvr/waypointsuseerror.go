// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type WaypointsUseError struct{}

func (m WaypointsUseError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.WaypointsUseError
}

func (m WaypointsUseError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsUseError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
