// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type WaypointsLeave struct{}

func NewWaypointsLeave(extra string) (WaypointsLeave, error) {
	var m WaypointsLeave

	if err := m.Deserialize(extra); err != nil {
		return WaypointsLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m WaypointsLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.WaypointsLeave
}

func (m WaypointsLeave) MessageName() string {
	return "WaypointsLeave"
}

func (m WaypointsLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
