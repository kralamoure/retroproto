// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommand struct{}

func NewBasicsAuthorizedCommand(extra string) (BasicsAuthorizedCommand, error) {
	var m BasicsAuthorizedCommand

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedCommand{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedCommand) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsAuthorizedCommand
}

func (m BasicsAuthorizedCommand) MessageName() string {
	return "BasicsAuthorizedCommand"
}

func (m BasicsAuthorizedCommand) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommand) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
