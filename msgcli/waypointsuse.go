// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type WaypointsUse struct{}

func (m WaypointsUse) ProtocolId() d1proto.MsgCliId {
	return d1proto.WaypointsUse
}

func (m WaypointsUse) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *WaypointsUse) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
