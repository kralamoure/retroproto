// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type WaypointsRequestLeave struct{}

func NewWaypointsRequestLeave(extra string) (WaypointsRequestLeave, error) {
	var m WaypointsRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return WaypointsRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m WaypointsRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.WaypointsRequestLeave
}

func (m WaypointsRequestLeave) MessageName() string {
	return "WaypointsRequestLeave"
}

func (m WaypointsRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
