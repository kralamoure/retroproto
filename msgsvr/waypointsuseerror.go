// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type WaypointsUseError struct{}

func NewWaypointsUseError(extra string) (WaypointsUseError, error) {
	var m WaypointsUseError

	if err := m.Deserialize(extra); err != nil {
		return WaypointsUseError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m WaypointsUseError) MessageId() retroproto.MsgSvrId {
	return retroproto.WaypointsUseError
}

func (m WaypointsUseError) MessageName() string {
	return "WaypointsUseError"
}

func (m WaypointsUseError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *WaypointsUseError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
