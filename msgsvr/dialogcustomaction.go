// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DialogCustomAction struct{}

func NewDialogCustomAction(extra string) (DialogCustomAction, error) {
	var m DialogCustomAction

	if err := m.Deserialize(extra); err != nil {
		return DialogCustomAction{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DialogCustomAction) MessageId() retroproto.MsgSvrId {
	return retroproto.DialogCustomAction
}

func (m DialogCustomAction) MessageName() string {
	return "DialogCustomAction"
}

func (m DialogCustomAction) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DialogCustomAction) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
