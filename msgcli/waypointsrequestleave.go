// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type WaypointsRequestLeave struct{}

func (m WaypointsRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.WaypointsRequestLeave
}

func (m WaypointsRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *WaypointsRequestLeave) Deserialize(extra string) error {
	return nil
}
