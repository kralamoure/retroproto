// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type WaypointsCreate struct{}

func (m WaypointsCreate) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.WaypointsCreate
}

func (m WaypointsCreate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *WaypointsCreate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
