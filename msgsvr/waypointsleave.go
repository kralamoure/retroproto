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
	return "", nil
}

func (m *WaypointsLeave) Deserialize(extra string) error {
	return nil
}
