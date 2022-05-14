// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedMoveCommand struct{}

func NewBasicsAuthorizedMoveCommand(extra string) (BasicsAuthorizedMoveCommand, error) {
	var m BasicsAuthorizedMoveCommand

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedMoveCommand{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedMoveCommand) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsAuthorizedMoveCommand
}

func (m BasicsAuthorizedMoveCommand) MessageName() string {
	return "BasicsAuthorizedMoveCommand"
}

func (m BasicsAuthorizedMoveCommand) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedMoveCommand) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
