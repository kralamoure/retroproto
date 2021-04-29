// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type WaypointsUse struct{}

func (m WaypointsUse) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.WaypointsUse
}

func (m WaypointsUse) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *WaypointsUse) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
