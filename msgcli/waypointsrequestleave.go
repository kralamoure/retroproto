// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type WaypointsRequestLeave struct{}

func (m WaypointsRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.WaypointsRequestLeave
}

func (m WaypointsRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *WaypointsRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
