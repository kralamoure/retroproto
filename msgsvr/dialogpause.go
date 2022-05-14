package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DialogPause struct{}

func NewDialogPause(extra string) (DialogPause, error) {
	var m DialogPause

	if err := m.Deserialize(extra); err != nil {
		return DialogPause{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DialogPause) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogPause
}

func (m DialogPause) MessageName() string {
	return "DialogPause"
}

func (m DialogPause) Serialized() (string, error) {
	return "", nil
}

func (m *DialogPause) Deserialize(extra string) error {
	return nil
}
