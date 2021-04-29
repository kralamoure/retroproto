// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type WaypointsLeave struct{}

func (m WaypointsLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.WaypointsLeave
}

func (m WaypointsLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *WaypointsLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
