// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedMoveCommand struct{}

func (m BasicsAuthorizedMoveCommand) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsAuthorizedMoveCommand
}

func (m BasicsAuthorizedMoveCommand) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsAuthorizedMoveCommand) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
