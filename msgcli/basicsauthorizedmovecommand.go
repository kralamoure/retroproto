// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedMoveCommand struct{}

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
