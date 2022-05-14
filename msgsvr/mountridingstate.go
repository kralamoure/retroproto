package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type MountRidingState struct {
	Riding bool
}

func (m MountRidingState) MessageId() retroproto.MsgSvrId {
	return retroproto.MountRidingState
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
