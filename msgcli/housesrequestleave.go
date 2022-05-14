// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesRequestLeave struct{}

func NewHousesRequestLeave(extra string) (HousesRequestLeave, error) {
	var m HousesRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return HousesRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.HousesRequestLeave
}

func (m HousesRequestLeave) MessageName() string {
	return "HousesRequestLeave"
}

func (m HousesRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
