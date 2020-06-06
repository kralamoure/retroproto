// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedKickCommand struct{}

func (m BasicsAuthorizedKickCommand) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsAuthorizedKickCommand
}

func (m BasicsAuthorizedKickCommand) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedKickCommand) Deserialize(extra string) error {
	return nil
}
