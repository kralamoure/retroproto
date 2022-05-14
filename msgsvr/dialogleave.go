package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DialogLeave struct{}

func NewDialogLeave(extra string) (DialogLeave, error) {
	var m DialogLeave

	if err := m.Deserialize(extra); err != nil {
		return DialogLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DialogLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogLeave
}

func (m DialogLeave) MessageName() string {
	return "DialogLeave"
}

func (m DialogLeave) Serialized() (string, error) {
	return "", nil
}

func (m *DialogLeave) Deserialize(extra string) error {
	return nil
}
