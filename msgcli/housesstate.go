// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesState struct{}

func NewHousesState(extra string) (HousesState, error) {
	var m HousesState

	if err := m.Deserialize(extra); err != nil {
		return HousesState{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesState) MessageId() retroproto.MsgCliId {
	return retroproto.HousesState
}

func (m HousesState) MessageName() string {
	return "HousesState"
}

func (m HousesState) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesState) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
