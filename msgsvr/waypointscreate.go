// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type WaypointsCreate struct{}

func NewWaypointsCreate(extra string) (WaypointsCreate, error) {
	var m WaypointsCreate

	if err := m.Deserialize(extra); err != nil {
		return WaypointsCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m WaypointsCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.WaypointsCreate
}

func (m WaypointsCreate) MessageName() string {
	return "WaypointsCreate"
}

func (m WaypointsCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
