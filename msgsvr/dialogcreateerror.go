package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DialogCreateError struct{}

func NewDialogCreateError(extra string) (DialogCreateError, error) {
	var m DialogCreateError

	if err := m.Deserialize(extra); err != nil {
		return DialogCreateError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DialogCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogCreateError
}

func (m DialogCreateError) MessageName() string {
	return "DialogCreateError"
}

func (m DialogCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *DialogCreateError) Deserialize(extra string) error {
	return nil
}
