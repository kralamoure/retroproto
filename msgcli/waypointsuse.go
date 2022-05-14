// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type WaypointsUse struct{}

func NewWaypointsUse(extra string) (WaypointsUse, error) {
	var m WaypointsUse

	if err := m.Deserialize(extra); err != nil {
		return WaypointsUse{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m WaypointsUse) MessageId() retroproto.MsgCliId {
	return retroproto.WaypointsUse
}

func (m WaypointsUse) MessageName() string {
	return "WaypointsUse"
}

func (m WaypointsUse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsUse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
