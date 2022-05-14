// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type WaypointsRequestLeave struct{}

func (m WaypointsRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.WaypointsRequestLeave
}

func (m WaypointsRequestLeave) MessageName() string {
	return "WaypointsRequestLeave"
}

func (m WaypointsRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
