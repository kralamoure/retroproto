// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedKickCommand struct{}

func (m BasicsAuthorizedKickCommand) ProtocolId() retroproto.MsgCliId {
	return retroproto.BasicsAuthorizedKickCommand
}

func (m BasicsAuthorizedKickCommand) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedKickCommand) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
