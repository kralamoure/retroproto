// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type WaypointsCreate struct{}

func (m WaypointsCreate) ProtocolId() d1proto.MsgSvrId {
	return d1proto.WaypointsCreate
}

func (m WaypointsCreate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *WaypointsCreate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
