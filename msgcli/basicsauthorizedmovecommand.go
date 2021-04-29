// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedMoveCommand struct{}

func (m BasicsAuthorizedMoveCommand) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsAuthorizedMoveCommand
}

func (m BasicsAuthorizedMoveCommand) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedMoveCommand) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
