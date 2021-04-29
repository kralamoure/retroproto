// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type WaypointsUseError struct{}

func (m WaypointsUseError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.WaypointsUseError
}

func (m WaypointsUseError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *WaypointsUseError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
