// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type WaypointsUse struct{}

func (m WaypointsUse) MessageId() retroproto.MsgCliId {
	return retroproto.WaypointsUse
}

func (m WaypointsUse) MessageName() string {
	return "WaypointsUse"
}

func (m WaypointsUse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsUse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
