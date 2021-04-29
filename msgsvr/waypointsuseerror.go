// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type WaypointsUseError struct{}

func (m WaypointsUseError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.WaypointsUseError
}

func (m WaypointsUseError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *WaypointsUseError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
