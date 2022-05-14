package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DialogRequestLeave struct{}

func NewDialogRequestLeave(extra string) (DialogRequestLeave, error) {
	var m DialogRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return DialogRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DialogRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.DialogRequestLeave
}

func (m DialogRequestLeave) MessageName() string {
	return "DialogRequestLeave"
}

func (m DialogRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogRequestLeave) Deserialize(extra string) error {
	return nil
}
