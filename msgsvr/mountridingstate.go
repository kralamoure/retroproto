package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountRidingState struct {
	Riding bool
}

func (m MountRidingState) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountRidingState
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
		return d1proto.ErrInvalidMsg
	}

	return nil
}
