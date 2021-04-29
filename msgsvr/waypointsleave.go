// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type WaypointsLeave struct{}

func (m WaypointsLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.WaypointsLeave
}

func (m WaypointsLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *WaypointsLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
