// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type WaypointsCreate struct{}

func (m WaypointsCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.WaypointsCreate
}

func (m WaypointsCreate) MessageName() string {
	return "WaypointsCreate"
}

func (m WaypointsCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
