package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type MountRidingState struct {
	Riding bool
}

func NewMountRidingState(extra string) (MountRidingState, error) {
	var m MountRidingState

	if err := m.Deserialize(extra); err != nil {
		return MountRidingState{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountRidingState) MessageId() retroproto.MsgSvrId {
	return retroproto.MountRidingState
}

func (m MountRidingState) MessageName() string {
	return "MountRidingState"
}

func (m MountRidingState) Serialized() (string, error) {
	var riding string
	if m.Riding {
		riding = "+"
	} else {
		riding = "-"
	}

	return riding, nil
}

func (m *MountRidingState) Deserialize(extra string) error {
	switch extra {
	case "+":
		m.Riding = true
	case "-":
		m.Riding = false
	default:
		return retroproto.ErrInvalidMsg
	}

	return nil
}
