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
	return "", nil
}

func (m *WaypointsUse) Deserialize(extra string) error {
	return nil
}
