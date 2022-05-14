// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedKickCommand struct{}

func (m BasicsAuthorizedKickCommand) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsAuthorizedKickCommand
}

func (m BasicsAuthorizedKickCommand) MessageName() string {
	return "BasicsAuthorizedKickCommand"
}

func (m BasicsAuthorizedKickCommand) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedKickCommand) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
