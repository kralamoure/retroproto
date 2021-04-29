// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedKickCommand struct{}

func (m BasicsAuthorizedKickCommand) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsAuthorizedKickCommand
}

func (m BasicsAuthorizedKickCommand) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedKickCommand) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
